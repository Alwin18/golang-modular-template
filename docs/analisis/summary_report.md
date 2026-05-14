# Code Review Report — golang-module-template

> **Reviewer**: Senior Go Engineer (Automated Review)
> **Date**: 2026-05-14
> **Scope**: Architecture, Security, Performance, Code Duplication
> **Codebase**: ~35 Go source files, Fiber v2 + GORM + Redis + Asynq

---

## 1. Architecture Summary

### Pattern: Modular Layered Architecture (Clean Architecture–lite)

```
cmd/api/main.go                    ← Entrypoint
    ↓
config/                            ← Config struct + .env loader
    ↓
internal/app/                      ← Bootstrap, DI Container, Worker registration
    ↓
internal/http/                     ← Fiber router + global middleware
    ↓
internal/module/{auth,health}/     ← Feature modules (handler → service → repository)
    ↓
internal/shared/                   ← Cross-cutting: cache, db, errors, response, security, validation
internal/worker/                   ← Async task workers (Asynq)
```

### Request Lifecycle

```
HTTP Request
  → Fiber Router (internal/http/route.go)
    → Global Middleware (recover → security headers → logger → CORS)
      → Route-specific Middleware (rate limit, auth JWT)
        → Handler (parse body, validate, call service)
          → Service (business logic, JWT generation)
            → Repository (GORM queries)
              → PostgreSQL
```

### Dependency Injection

Manual struct-based container (`internal/app/container.go`). A flat `Deps` struct is passed to the router to avoid import cycles. Modules are wired explicitly in `RegisterRoutes()`.

### What's Good

- Clear module boundary pattern with `domain.go`, `handler.go`, `service.go`, `repository.go`, `route.go`
- Good separation of shared concerns into `internal/shared/`
- Redis-based rate limiting on auth endpoints
- Security headers middleware
- Typed error system with `AppError` + status code mapping
- Proper bcrypt cost (12) for password hashing
- Connection pool tuning on Postgres
- `BlacklistTokenKey` and `LockKey` helpers ready for future use
- Sample module as a template for new features

---

## 2. Problem Areas

| # | File/Package | Issue | Severity | Why It Matters |
|---|---|---|---|---|
| 1 | `internal/http/route.go:46` | CORS allows all origins (`cors.New()` default) | 🔴 Critical | Any website can make authenticated requests to your API |
| 2 | `internal/shared/errors/error.go:110-129` | Error status mapping via hardcoded Indonesian string matching | 🔴 Critical | Any typo or translation change silently breaks HTTP status codes |
| 3 | `internal/http/route.go:80-89` | `defaultErrorHandler` leaks raw error messages to clients | 🔴 Critical | Internal errors/stack traces exposed in production |
| 4 | `.env` committed + `JWT_SECRET=supersecret` | Secrets in version control | 🔴 Critical | Credential leak; validation only blocks in production mode |
| 5 | `internal/http/middleware/auth.go` | No token blacklist/revocation check | 🔴 Critical | Logged-out tokens remain valid until expiry |
| 6 | `internal/worker/webhook_worker.go:56` | No URL validation — SSRF risk | 🔴 Critical | Server can be tricked into hitting internal services |
| 7 | `internal/worker/webhook_worker.go:40` | Context parameter ignored — no cancellation | 🟡 Medium | Tasks can't be cancelled; no timeout propagation |
| 8 | `internal/module/auth/repository.go:19` | `Login()` accepts unused `password` param | 🟡 Medium | Misleading API; the repo is really `FindByUsername` |
| 9 | `internal/module/auth/service.go:36,65` | Direct `user.Role.Name` access with no nil check | 🟡 Medium | Panic if Role preload fails or role is nil |
| 10 | All repository methods | No `context.Context` parameter | 🟡 Medium | No query timeouts, no request cancellation |
| 11 | `internal/shared/slotpool/pool.go:33` | Pipeline created but never executed (`Exec()` not called) | 🟡 Medium | `Acquire()` silently does nothing — broken feature |
| 12 | `internal/shared/cache/cache.go:56-62` | Rate limit TTL resets on every increment (sliding window) | 🟡 Medium | Attackers can space requests to never hit the limit |
| 13 | `internal/shared/errors/error.go:160-162` | Global mutable map without sync | 🟡 Medium | Race condition if called concurrently |
| 14 | `internal/module/health/handler.go:49` | Response message: "success login" (copy-paste) | 🟢 Minor | Wrong message in health check response |
| 15 | `response.go` + `utils.go` | Duplicate `TotalPage` / `Pagination` logic | 🟢 Minor | Maintenance burden, inconsistency risk |
| 16 | `internal/shared/response/response.go:53-58` | `NewErrorResponse` is a no-op identity function | 🟢 Minor | Unnecessary abstraction |
| 17 | `internal/shared/db/postgres.go:14` | GORM logger hardcoded to `Silent` | 🟢 Minor | No SQL debugging possible |
| 18 | `internal/shared/db/models/user.go:9` | `Username` has no unique index | 🟢 Minor | Duplicate usernames possible; login returns wrong user |
| 19 | `internal/logging/`, `middleware/permission.go` | Empty packages (dead code) | 🟢 Minor | Confusing for new developers |
| 20 | `internal/worker/queue.go` | Task constructors always return nil error | 🟢 Minor | Misleading function signature |

---

## 3. Refactoring Strategies

### 🔴 Issue #1 — CORS Allows All Origins

**Problem**: `cors.New()` with no configuration defaults to `AllowOrigins: "*"`. Any malicious website can make authenticated cross-origin requests.

**Before** (`internal/http/route.go:46`):
```go
fiberApp.Use(cors.New())
```

**After**:
```go
fiberApp.Use(cors.New(cors.Config{
    AllowOrigins:     deps.Config.CORSOrigins, // e.g. "https://app.example.com"
    AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
    AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
    AllowCredentials: true,
    MaxAge:           3600,
}))
```

Add `CORSOrigins string` to `config.Config` and load from `CORS_ORIGINS` env var.

---

### 🔴 Issue #2 — Fragile String-Based Error Mapping

**Problem**: `GetStatusCode()` falls back to matching Indonesian error message strings. Any change to the message text silently returns 500 instead of the correct status.

**Before** (`internal/shared/errors/error.go:109-129`):
```go
// Fallback: check by error message for backward compatibility
errMsg := err.Error()
switch errMsg {
case "data tidak ditemukan",
    "akun tidak ditemukan",
    // ...
    return fiber.StatusNotFound
// ...
```

**After** — Remove the entire string-matching block. The `errorStatusMap` already covers all constants:
```go
func GetStatusCode(err error) int {
    if appErr, ok := err.(*AppError); ok {
        return appErr.StatusCode
    }

    for knownErr, statusCode := range errorStatusMap {
        if errors.Is(err, knownErr) {
            return statusCode
        }
    }

    return fiber.StatusInternalServerError
}
```

---

### 🔴 Issue #3 — Raw Error Messages Leaked to Clients

**Problem**: `defaultErrorHandler` sends `err.Error()` directly. In production, Go internal errors (DB connection strings, file paths, panic messages) get exposed.

**Before** (`internal/http/route.go:80-89`):
```go
func defaultErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError
    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }
    return c.Status(code).JSON(fiber.Map{
        "success": false,
        "message": err.Error(), // ← leaks internals
    })
}
```

**After**:
```go
func defaultErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError
    message := "internal server error"

    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
        message = e.Message
    }
    if appErr, ok := err.(*apperrors.AppError); ok {
        code = appErr.StatusCode
        message = appErr.Message
    }

    return c.Status(code).JSON(fiber.Map{
        "success": false,
        "message": message,
    })
}
```

---

### 🔴 Issue #4 — Secrets in Version Control

**Problem**: `.env` is committed with `JWT_SECRET=supersecret` and database credentials.

**Fix**: Add `.env` to `.gitignore` (keep `.env.example` only):
```gitignore
# current .gitignore content + add:
.env
```

---

### 🔴 Issue #5 — No Token Blacklist Check

**Problem**: After logout, the JWT remains valid until expiry. `BlacklistTokenKey()` exists but is never used.

**Before** (`internal/http/middleware/auth.go:30-36`):
```go
claims, err := jwtManager.ParseToken(parts[1])
if err != nil {
    return ctx.Status(fiber.StatusUnauthorized).JSON(...)
}
```

**After** — Add blacklist check (requires passing `cache` to the middleware):
```go
func Auth(jwtManager *security.JWTManager, cacheClient *cache.Cache) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        // ... existing header parsing ...

        claims, err := jwtManager.ParseToken(parts[1])
        if err != nil {
            return ctx.Status(fiber.StatusUnauthorized).JSON(...)
        }

        // Check token blacklist
        _, err = cacheClient.Get(ctx.Context(), cache.BlacklistTokenKey(parts[1]))
        if err == nil {
            // Token found in blacklist
            return ctx.Status(fiber.StatusUnauthorized).JSON(response.NewErrorResponse(response.ResponseError{
                Message: "token has been revoked",
                Code:    fiber.StatusUnauthorized,
            }))
        }

        ctx.Locals("user_id", claims.UserID)
        // ...
    }
}
```

---

### 🔴 Issue #6 — Webhook SSRF

**Problem**: `WebhookPayload.URL` is used in `http.NewRequest` with no validation. Internal URLs (e.g., `http://169.254.169.254/`) could be targeted.

**After** — Add URL validation:
```go
func (w *WebhookWorker) ProcessTask(ctx context.Context, t *asynq.Task) error {
    // ... unmarshal ...

    parsed, err := url.Parse(payload.URL)
    if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
        return fmt.Errorf("invalid webhook URL: %w", asynq.SkipRetry)
    }

    // Block private/internal IPs
    ips, err := net.LookupIP(parsed.Hostname())
    if err != nil {
        return fmt.Errorf("dns lookup failed: %w", err)
    }
    for _, ip := range ips {
        if ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() {
            return fmt.Errorf("webhook to internal IP blocked: %w", asynq.SkipRetry)
        }
    }

    // Use context for cancellation
    req, err := http.NewRequestWithContext(ctx, method, payload.URL, bytes.NewBuffer(body))
    // ...
}
```

---

### 🟡 Issue #8 — Misleading Repository Method

**Before** (`internal/module/auth/repository.go:19`):
```go
func (r *Repository) Login(username, password string) (models.User, error) {
```

**After**:
```go
func (r *Repository) FindByUsername(username string) (models.User, error) {
```

Update the caller in `service.go:27`:
```go
user, err := s.repo.FindByUsername(req.Username)
```

---

### 🟡 Issue #10 — No Context in Repository

**Before**:
```go
func (r *Repository) Login(username, password string) (models.User, error) {
    var user models.User
    if err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
```

**After**:
```go
func (r *Repository) FindByUsername(ctx context.Context, username string) (models.User, error) {
    var user models.User
    if err := r.db.WithContext(ctx).Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
```

Apply the same pattern to `SaveRefreshToken` and `DeleteRefreshToken`. Pass `ctx` from handler → service → repository.

---

### 🟡 Issue #11 — SlotPool Pipeline Bug

**Problem**: `p.client.Pipeline().SetNX(...)` creates a pipeline but never calls `Exec()`. The `SetNX` is never sent to Redis.

**Before** (`internal/shared/slotpool/pool.go:33`):
```go
ok, err := p.client.Pipeline().SetNX(ctx, key, "1", ttl).Result()
```

**After**:
```go
ok, err := p.client.SetNX(ctx, key, "1", ttl).Result()
```

---

### 🟡 Issue #12 — Sliding Window Rate Limit

**Problem**: `Increment()` always calls `Expire()`, resetting the TTL on every request. This creates a sliding window instead of a fixed window.

**After** — Only set expiry on first increment:
```go
func (c *Cache) Increment(ctx context.Context, key string, ttl time.Duration) (int64, error) {
    count, err := c.client.Incr(ctx, key).Result()
    if err != nil {
        return 0, err
    }
    // Only set TTL on first increment (when count becomes 1)
    if count == 1 {
        c.client.Expire(ctx, key, ttl)
    }
    return count, nil
}
```

---

## 4. Additional Recommendations

### Missing Implementations (per README)

The README documents these routes but they are **not implemented**:

| Route | Status |
|---|---|
| `POST /api/v1/auth/register` | ❌ Missing |
| `POST /api/v1/auth/refresh` | ❌ Missing |
| `POST /api/v1/auth/logout` | ❌ Missing |
| `GET /api/v1/auth/me` | ❌ Missing |
| `GET /api/v1/users` | ❌ Missing (entire user module) |
| `GET /api/v1/users/:id` | ❌ Missing |
| `PUT /api/v1/users/:id` | ❌ Missing |
| `DELETE /api/v1/users/:id` | ❌ Missing |

### Quick Wins

1. **Add `.env` to `.gitignore`** — one-line fix, immediate security improvement
2. **Fix health handler message** — change `"success login"` to `"health check"` (line 49)
3. **Add unique index on `users.username`** — `gorm:"uniqueIndex"` tag
4. **Delete empty packages** — `internal/logging/worker.go`, `internal/http/middleware/permission.go`
5. **Make GORM log level configurable** — use `logger.Info` in development, `logger.Silent` in production

### Testing

No test files exist in the codebase. Consider adding:
- Unit tests for `security/` (JWT generation/parsing, password hashing)
- Unit tests for `errors/` (status code mapping)
- Integration tests for auth flow
- Table-driven tests for validation logic

---

## 5. Priority Action Plan

| Priority | Action | Effort |
|---|---|---|
| 🔴 P0 | Configure CORS with explicit allowed origins | 10 min |
| 🔴 P0 | Add `.env` to `.gitignore`, rotate JWT secret | 5 min |
| 🔴 P0 | Sanitize error messages in `defaultErrorHandler` | 15 min |
| 🔴 P0 | Remove string-matching fallback in `GetStatusCode` | 5 min |
| 🔴 P0 | Add token blacklist check in auth middleware | 30 min |
| 🔴 P0 | Add URL validation to webhook worker | 30 min |
| 🟡 P1 | Add `context.Context` to all repository methods | 1 hr |
| 🟡 P1 | Fix SlotPool pipeline bug | 5 min |
| 🟡 P1 | Fix rate limiter sliding window | 10 min |
| 🟡 P1 | Rename `repo.Login` → `repo.FindByUsername` | 10 min |
| 🟡 P1 | Add nil guard for `user.Role` | 5 min |
| 🟢 P2 | Clean up empty packages, fix typos | 15 min |
| 🟢 P2 | Add unit tests for shared packages | 2-4 hrs |
| 🟢 P2 | Add username unique index | 5 min |

---

*End of Report*
