# Сокращатель ссылок RESTAPI

### Архитектура

```bash
├── cmd
│   └── url-shortener
│       └── main.go          
├── internal
│   ├── http-server          
│   │   └── handlers
│   │       ├── redirect     
│   │       │   └── redirect.go
│   │       └── shortener    
│   │           └── shortener.go
│   ├── domain               
│   │   └── link.go          
│   ├── services             
│   │   └── services.go     
│   ├── repository           
│   │   └── link.go                  - методы к базе данных
│   ├── config                       - конфигурационные файлы
│   │   ├── app.go
│   │   ├── database.go
│   │   └── server.go
│   └── database             
│       └── database.go      
├── docker-compose.yml       
├── go.mod                   
├── go.sum                   
├── Makefile                 
└── samples                 
