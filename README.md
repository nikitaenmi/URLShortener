# Сокращатель ссылок REST API

### Архитектура

```bash
├── cmd
│   └── url-shortener
│       └── main.go
├── internal
│   ├── http-server
│   │   ├── handlers
│   │   │   └── url.go          - HTTP-обработчики
│   │   └── middleware
│   │       └── request_id.go
│   ├── domain
│   │   └── url.go
│   ├── services
│   │   └── url.go              - бизнес-логика
│   ├── repository
│   │   └── url.go              - слой работы с данными
│   ├── config
│   │   ├── app.go
│   │   ├── database.go
│   │   ├── generator.go        - генерация случайного алиаса
│   │   └── server.go
│   ├── database
│   │   └── database.go
│   └── lib
│       ├── generator
│       │   ├── digit.go
│       │   ├── init.go
│       │   └── lowercase.go
│       └── logger
│           ├── logger.go
│           └── slog
│               └── handler
│                   └── ctx.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── samples
