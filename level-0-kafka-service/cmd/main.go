package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/config"
	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/handler"
	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/kafka"
	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/repository"
	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := loadConfig()
	logger := initLogger()

	db := connectDatabase(cfg, logger)
	defer db.Close()

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	router := setupRouter(orderHandler)

	runKafka(cfg, orderService)
	startServer(cfg, router, logger)
}

func runKafka(cfg *config.Config, orderService *service.OrderService) {
	consumer := kafka.NewConsumer(
		cfg.Kafka.Broker,
		cfg.Kafka.Topic,
		cfg.Kafka.GroupID,
		orderService,
	)

	go consumer.Run(context.Background())
}

func loadConfig() *config.Config {
	cfg, err := config.LoadConfigFromEnv()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}
	return cfg
}

func initLogger() *slog.Logger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)
	return logger
}

func connectDatabase(cfg *config.Config, logger *slog.Logger) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.SQL.Username,
		cfg.SQL.Password,
		cfg.SQL.Host,
		cfg.SQL.Port,
		cfg.SQL.Name,
	)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Error("failed to connect to PostgreSQL", "err", err)
		os.Exit(1)
	}
	return dbpool
}

func setupRouter(orderHandler *handler.OrderHandler) http.Handler {
	r := chi.NewRouter()
	r.Route("/order", func(r chi.Router) {
        r.Get("/{uid}", orderHandler.GetByID) 
		r.Get("/health", orderHandler.Health)
    })
	return r
}

func startServer(cfg *config.Config, handler http.Handler, logger *slog.Logger) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		logger.Info("starting HTTP server", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server failed", "err", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("graceful shutdown failed", "err", err)
	}
}
