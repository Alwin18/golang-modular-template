# Golang Modular Template

A production-ready **Golang backend template** using **Modular + Layered Architecture**, designed for scalability, maintainability, and clarity.

This template is suitable for monolith or growing systems that may evolve into microservices.

---

## ✨ Features

- Modular & feature-based structure
- Layered architecture (Handler → Service → Repository)
- Manual Dependency Injection (no magic)
- Fiber HTTP framework
- GORM ORM (multi-database ready)
- Redis integration
- Logrus logging
- Validator for request validation
- Clean separation between:
  - Application bootstrap
  - Business modules
  - Infrastructure
  - Shared utilities

---

## 🧱 Architecture Overview

### Folder Structure

```
├─ cmd/
│ └─ api/
│ ├─ main.go # Application entry point
│ └─ config.go # Load env / config
│
├─ internal/
│ ├─ app/
│ │ ├─ container.go # Dependency wiring (DB, redis, logger, services)
│ │ └─ app.go # Fiber bootstrap & global middleware
│ │
│ ├─ module/
│ │ ├─ user/
│ │ │ ├─ handler.go # Fiber HTTP handler
│ │ │ ├─ service.go # Business logic
│ │ │ ├─ repository.go # Repository interface
│ │ │ ├─ repository_gorm.go # GORM implementation
│ │ │ ├─ model.go # Entity / DTO
│ │ │ └─ route.go # Route registration
│ │ │
│ │ ├─ order/
│ │ ├─ handler.go
│ │ ├─ service.go
│ │ ├─ repository.go
│ │ ├─ repository_gorm.go
│ │ ├─ model.go
│ │ └─ route.go
│ │
│ ├─ http/
│ │ ├─ router.go # Global route registration
│ │ └─ middleware/
│ │ ├─ auth.go # Auth / JWT middleware
│ │ ├─ logging.go # Request logging
│ │ └─ recover.go # Panic recovery
│ │
│ ├─ shared/
│ │ ├─ db/ # GORM initialization & DB helpers
│ │ ├─ redis/ # Redis client & cache abstraction
│ │ ├─ logger/ # Logrus setup
│ │ ├─ validation/ # Validator helper
│ │ ├─ pagination/ # Pagination helper
│ │ ├─ crypto/ # Password hashing / crypto helpers
│ │ └─ errors/ # Application & HTTP error mapping
│
├─ go.mod
└─ go.sum
```


---

## 🔄 Application Flow

```
main.go
└─ LoadConfig()
└─ app.NewContainer()
├─ initialize logger
├─ initialize database(s)
├─ initialize redis
└─ build services
└─ inject repositories & dependencies
↓
└─ app.NewApp()
├─ fiber.New()
├─ register global middleware
└─ http.RegisterRoutes()
└─ module.RegisterRoutes()
```


---

## 🧠 Architectural Principles

### Modular
- Each feature lives in its own module
- No cross-module imports
- Modules communicate only through injected dependencies

### Layered (inside module)

- **Handler**: HTTP concerns only
- **Service**: Business logic
- **Repository**: Data access

---

## 🧩 Dependency Injection

- All dependencies are wired in `internal/app/container.go`
- No global variables
- No framework-based DI
- Explicit, testable, and Go-idiomatic

---

## 🗄 Database & Redis

### GORM
- Initialized in `shared/db`
- Supports multiple databases
- Injected into repositories

### Redis
- Initialized in `shared/redis`
- Can be used for caching, session, or pub/sub
- Injected via service or repository

---

## 🧪 Validation

- Uses `go-playground/validator`
- Centralized in `shared/validation`
- Used mainly in HTTP handlers

---

## 📜 Logging

- Uses `logrus`
- Centralized configuration in `shared/logger`
- Logger injected into services or middleware

---

## 🚦 Error Handling

- Domain/application errors live in `shared/errors`
- HTTP mapping handled centrally
- Services return domain errors, not HTTP errors

---

## 🧰 How to Add a New Module

1. Create a new folder under `internal/module/`
2. Add:
   - `handler.go`
   - `service.go`
   - `repository.go`
   - `repository_gorm.go`
   - `model.go`
   - `route.go`
3. Wire the module in `app/container.go`
4. Register routes in `http/router.go`

---

## 🧱 Scaling the Project

This structure supports:
- Large monoliths
- Multiple databases
- Background workers (`cmd/worker`)
- Cron jobs (`cmd/cron`)
- Migration tools (`cmd/migrate`)
- Gradual evolution to Clean / Hexagonal Architecture

---

## ✅ When to Use This Template

- REST API backends
- CRUD-heavy systems
- Enterprise applications
- Long-lived projects
- Teams that value clarity over magic

---

## 📌 License

MIT (or your preferred license)

---

## 🙌 Final Notes

This template prioritizes:
- Explicit dependencies
- Clear boundaries
- Practical scalability
- Real-world Go conventions

Feel free to fork, adapt, and evolve it to your needs.
