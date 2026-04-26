package sample

// Sample module — copy this pattern when creating a new module.
//
// Steps:
//  1. Define your domain types in domain.go
//  2. Implement data access in repository.go  (depends on *gorm.DB)
//  3. Implement business logic in service.go  (depends on *Repository)
//  4. Expose HTTP handlers in handler.go      (depends on *Service)
//  5. Register routes in route.go             (called from internal/http/route.go)
//
// Flow: HTTP → Handler → Service → Repository → DB
