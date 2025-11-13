# Сокращатель ссылок REST API


## Содержание

- [Архитектура](#архитектура)


## Архитектура 

```bash
url-shortener/
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── docker-compose.yml
├── cmd/
│   └── url-shortener/
│       └── main.go                - точка входа
├── internal/
│   ├── domain/                    - доменные сущности
│   │   └── url.go
│   ├── config/                    - конфиг
│   │   ├── app.go               
│   │   ├── server.go
│   │   ├── database.go
│   │   └── generator.go
│   ├── repository/                - репозиторный слой
│   │   └── url.go
│   ├── services/                  - сервисный слой
│   │   └── url.go
│   ├── http-server/            
│   │   ├── dto.go
│   │   ├── mapping.go
│   │   ├── handlers/              - хендлеры
│   │   │   ├── url.go
│   │   │   └── utils.go
│   │   └── middleware/
│   │       ├── error_handler.go   - обработчик ошибок HTTP
│   │       └── request_id.go      - мидлварь для Request ID
│   └── lib/
│       ├── logger/
│       │   ├── logger.go          - интерфейс логгера
│       │   ├── context_logger.go  - логгер с поддержкой контекста
│       │   └── echo_middleware.go - middleware для логирования запросов 
│       └── generator/
│           ├── init.go
│           ├── digit.go
│           └── lowercase.go
└── samples/.env.example           - пример .env файла
