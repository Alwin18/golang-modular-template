# Golang Modular Template

## Description

This is a modular template for Golang applications.

## Architecture
```
.
├─ cmd/
│  └─ api/
│     ├─ main.go                 # Entry point (super tipis)
│     └─ config.go               # Load env / config
│
├─ internal/
│  ├─ app/
│  │  ├─ container.go            # Dependency wiring (DB, redis, logger, services)
│  │  └─ app.go                  # Fiber bootstrap (new app, global middleware)
│  │
│  ├─ module/
│  │  ├─ user/
│  │  │  ├─ handler.go           # Fiber HTTP handler
│  │  │  ├─ service.go           # Business logic
│  │  │  ├─ repository.go        # Repository interface
│  │  │  ├─ repository_gorm.go   # GORM implementation
│  │  │  ├─ model.go             # Entity / DTO
│  │  │  └─ route.go             # Route register for module
│  │  │
│  │  ├─ order/
│  │  │  ├─ handler.go
│  │  │  ├─ service.go
│  │  │  ├─ repository.go
│  │  │  ├─ repository_gorm.go
│  │  │  ├─ model.go
│  │  │  └─ route.go
│  │
│  ├─ http/
│  │  ├─ router.go               # Call module route register
│  │  └─ middleware/
│  │     ├─ auth.go              # Auth / JWT middleware
│  │     ├─ logging.go           # Request logging
│  │     └─ recover.go           # Panic recovery
│  │
│  ├─ shared/
│  │  ├─ db/
│  │  │  ├─ gorm.go              # Init *gorm.DB
│  │  │  ├─ postgres.go          # Postgres dialector & options
│  │  │  └─ mysql.go             # MySQL dialector (optional)
│  │  │
│  │  ├─ redis/
│  │  │  ├─ client.go            # Redis client init
│  │  │  └─ cache.go             # Cache abstraction (Get/Set)
│  │  │
│  │  ├─ logger/
│  │  │  └─ logger.go            # Logrus setup
│  │  │
│  │  ├─ validation/
│  │  │  └─ validator.go         # Validator instance & helper
│  │  │
│  │  ├─ pagination/
│  │  │  └─ pagination.go        # Limit/offset helper
│  │  │
│  │  ├─ crypto/
│  │  │  └─ password.go          # Hash / compare password
│  │  │
│  │  └─ errors/
│  │     ├─ app_error.go         # Base app error
│  │     └─ http_error.go        # HTTP mapping helper
│
├─ go.mod
└─ go.sum
```

## Flow
```
main.go
 └─ LoadConfig()
     └─ app.NewContainer()
         ├─ init logger
         ├─ init gorm (multi DB ready)
         ├─ init redis
         └─ build services
             └─ inject repo, logger, cache
                 ↓
     └─ app.NewApp()
         ├─ fiber.New()
         ├─ global middleware
         └─ http.RegisterRoutes()
             └─ module.RegisterRoutes()
```