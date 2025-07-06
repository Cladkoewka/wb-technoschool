# WB TechnoSchool Level 0 — Order Service

Микросервис, принимающий данные заказов через Kafka, сохраняющий их в PostgreSQL и предоставляющий HTTP-интерфейс для получения данных по `order_uid`. Сервис использует in-memory кэш для ускорения доступа и восстановления данных при перезапуске.


## 📁 Структура проекта

```text
.
├── cmd/
│   └── main.go                # Точка входа
├── internal/
│   ├── config/                # Загрузка конфигурации из env
│   ├── domain/                # Модели: Order, Delivery, Payment, Item
│   ├── handler/               # HTTP-обработчики
│   ├── kafka/                 # Kafka consumer
│   ├── repository/            # Работа с PostgreSQL
│   ├── service/               # Бизнес-логика и кэш
├── migrations/               # goose миграции
├── scripts/
│   └── producer.go           # Пример Kafka producer-а
├── web/
│   └── index.html            # HTML-интерфейс поиска по UID
├── .env                      # Конфигурация среды
├── README.md
```
