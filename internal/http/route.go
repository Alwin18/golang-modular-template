package http

import (
	"github.com/Alwin18/golang-module-template/config"
	"github.com/Alwin18/golang-module-template/internal/http/middleware"
	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	"github.com/Alwin18/golang-module-template/internal/shared/security"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	authModule "github.com/Alwin18/golang-module-template/internal/module/auth"
	healthModule "github.com/Alwin18/golang-module-template/internal/module/health"
)

// Deps holds the flat set of dependencies the router needs.
// This avoids importing internal/app and creating an import cycle.
type Deps struct {
	Config    *config.Config
	DB        *gorm.DB
	Redis     *redis.Client
	Cache     *cache.Cache
	Logger    *zap.Logger
	Validator *validator.Validate
}

// Router wraps Fiber and registers all routes.
type Router struct {
	app  *fiber.App
	deps *Deps
}

// NewRouter creates a new Router with global middleware.
func NewRouter(deps *Deps) *Router {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: defaultErrorHandler,
	})

	// Global middleware
	fiberApp.Use(middleware.Recover())
	fiberApp.Use(middleware.Security())
	fiberApp.Use(middleware.Logger(deps.Logger))
	fiberApp.Use(cors.New())

	return &Router{app: fiberApp, deps: deps}
}

// RegisterRoutes wires all module routes.
func (r *Router) RegisterRoutes() {
	jwtManager := security.NewJWTManager(
		r.deps.Config.JWTSecret,
		r.deps.Config.JWTAccessTokenTTL,
		r.deps.Config.JWTRefreshTokenTTL,
	)
	authMW := middleware.Auth(jwtManager)

	api := r.app.Group("/api/v1")

	// Health
	healthHandler := healthModule.NewHandler(r.deps.DB, r.deps.Redis)
	healthModule.RegisterRoutes(api, healthHandler)

	// Auth
	authRepo := authModule.NewRepository(r.deps.DB)
	authSvc := authModule.NewService(authRepo, jwtManager, r.deps.Cache)
	authHandler := authModule.NewHandler(authSvc, r.deps.Validator)
	authModule.RegisterRoutes(api, authHandler, r.deps.Cache, authMW)

}

// App returns the underlying Fiber app.
func (r *Router) App() *fiber.App {
	return r.app
}

// defaultErrorHandler handles unhandled errors.
func defaultErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"message": err.Error(),
	})
}
