# starter-service

Production-ready Go REST API skeleton — modular, clean, and scalable.

## Stack

| Layer      | Tech           |
|------------|----------------|
| HTTP       | Fiber v2       |
| ORM        | GORM           |
| Database   | PostgreSQL     |
| Cache      | Redis          |
| Auth       | JWT + bcrypt   |
| Queue      | Asynq          |
| Logging    | Zap            |
| Config     | godotenv       |

## Quick Start

```bash
cp .env.example .env
make docker-up   # start postgres + redis
make run         # start server on :8080
```

## API Endpoints

| Method | Path                    | Auth | Description       |
|--------|-------------------------|------|-------------------|
| POST   | /api/v1/auth/register   | ❌   | Register          |
| POST   | /api/v1/auth/login      | ❌   | Login             |
| POST   | /api/v1/auth/refresh    | ❌   | Refresh token     |
| POST   | /api/v1/auth/logout     | ✅   | Logout            |
| GET    | /api/v1/auth/me         | ✅   | Current user      |
| GET    | /api/v1/users           | ✅   | List users        |
| GET    | /api/v1/users/:id       | ✅   | Get user          |
| PUT    | /api/v1/users/:id       | ✅   | Update user       |
| DELETE | /api/v1/users/:id       | ✅   | Delete user       |
| GET    | /api/v1/health          | ❌   | Health check      |

## Folder Structure

```
starter-service/
├── cmd/
│   └── api/
│       └── main.go               # Entrypoint
│
├── config/
│   ├── config.go                 # Config struct
│   └── env.go                    # Load from .env
│
├── internal/
│   ├── app/
│   │   ├── app.go                # Bootstrap & graceful shutdown
│   │   ├── container.go          # Dependency injection container
│   │   └── workers.go            # Register Asynq workers
│   │
│   ├── http/
│   │   ├── route.go              # Router setup, Deps struct
│   │   └── middleware/
│   │       ├── auth.go           # JWT validation
│   │       ├── logging.go        # Structured request logging
│   │       ├── ratelimit.go      # Redis-based rate limiting
│   │       ├── recover.go        # Panic recovery
│   │       └── security.go       # Security headers
│   │
│   ├── module/
│   │   ├── auth/                 # Register, login, refresh, logout, me
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── repository.go
│   │   │   ├── route.go
│   │   │   └── service.go
│   │   ├── user/                 # CRUD users
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── repository.go
│   │   │   ├── route.go
│   │   │   └── service.go
│   │   ├── health/               # DB + Redis liveness check
│   │   │   └── handler.go
│   │   └── sample/               # Template for new modules
│   │
│   ├── worker/
│   │   ├── queue.go              # Task type constants
│   │   ├── email_worker.go       # Send email async
│   │   └── webhook_worker.go     # Deliver webhooks async
│   │
│   └── shared/
│       ├── cache/                # Redis wrapper (Get/Set/Remember/Increment)
│       ├── db/                   # Postgres connection + GORM models
│       ├── errors/               # Typed AppError with HTTP codes
│       ├── logger/               # Zap logger factory
│       ├── response/             # HTTP response helpers
│       ├── security/             # JWT manager, bcrypt, SHA256
│       ├── slotpool/             # Distributed lock (prevent duplicate ops)
│       ├── utils/                # Pagination, sanitize helpers
│       └── validation/           # Chainable field validator
│
│
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── go.mod
```

## Adding a New Module

1. Create `internal/module/yourmodule/` with: `domain.go`, `repository.go`, `service.go`, `handler.go`, `route.go`
2. Wire it in `internal/http/route.go`

Flow: `HTTP → Handler → Service → Repository → DB`

## Make Commands

```bash
make run            # run dev server
make build          # compile binary to bin/
make test           # run all tests
make docker-up      # start postgres + redis
make docker-down    # stop containers
```

## Environment Variables

```env
APP_PORT=8080
APP_ENV=local

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=secret
DB_NAME=starter

REDIS_HOST=localhost
REDIS_PORT=6379

JWT_SECRET=change_me_in_production
JWT_ACCESS_TOKEN_TTL=15    # minutes
JWT_REFRESH_TOKEN_TTL=7    # days
```