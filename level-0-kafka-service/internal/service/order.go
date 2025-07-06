package service

import (
	"context"
	"log/slog"
	"sync"

	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/domain"
)

type OrderRepository interface {
	SaveOrder(ctx context.Context, order *domain.Order) error
	GetOrderByID(ctx context.Context, orderUID string) (*domain.Order, error)
	GetOrders(ctx context.Context) ([]*domain.Order, error)
}

type OrderService struct {
	repo OrderRepository

	cacheMu sync.RWMutex
	cache   map[string]*domain.Order
}

func NewOrderService(repo OrderRepository) *OrderService {
	s := &OrderService{
		repo:  repo,
		cache: make(map[string]*domain.Order),
	}
	if err := s.loadCacheFromDB(context.Background()); err != nil {
		slog.Error("failed to load cache from DB", "err", err)
	} else {
		slog.Info("cache loaded from DB successfully", "items", len(s.cache))
	}
	return s
}

func (s *OrderService) ProcessOrder(ctx context.Context, order *domain.Order) error {
	if err := s.repo.SaveOrder(ctx, order); err != nil {
		slog.Error("failed to save order to DB", "order_uid", order.OrderUID, "err", err)
		return err
	}

	// s.cacheMu.Lock()
	// s.cache[order.OrderUID] = order
	// s.cacheMu.Unlock()
	// slog.Info("cache updated with order", "order_uid", order.OrderUID)

	return nil
}

func (s *OrderService) GetOrder(ctx context.Context, orderUID string) (*domain.Order, error) {
	s.cacheMu.RLock()
	order, ok := s.cache[orderUID]
	s.cacheMu.RUnlock()
	if ok {
		slog.Info("order found in cache", "order_uid", orderUID)
		return order, nil
	}

	order, err := s.repo.GetOrderByID(ctx, orderUID)
	if err != nil {
		slog.Error("failed to get order from DB", "order_uid", orderUID, "err", err)
		return nil, err
	}
	if order == nil {
		slog.Warn("order not found in DB", "order_uid", orderUID)
		return nil, nil
	}

	s.cacheMu.Lock()
	s.cache[orderUID] = order
	s.cacheMu.Unlock()
	slog.Info("order added to cache after DB fetch", "order_uid", orderUID)

	return order, nil
}

func (s *OrderService) loadCacheFromDB(ctx context.Context) error {
	orders, err := s.repo.GetOrders(ctx)
	if err != nil {
		return err
	}
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()
	for _, order := range orders {
		s.cache[order.OrderUID] = order
	}
	return nil
}
