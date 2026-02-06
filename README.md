# Golang Modular Template

## Description

This is a modular template for Golang applications.

## Architecture
```
internal/
 ├─ user/
 │   ├─ handler.go
 │   ├─ service.go
 │   ├─ repository.go
 │   └─ model.go
 │
 ├─ order/
 │   ├─ handler.go
 │   ├─ service.go
 │   ├─ repository.go
 │   └─ model.go
 │
 ├─ shared/
 │   ├─ db/
 │   ├─ logger/
 │   ├─ errors/
 |   |─ middleware/
 |   |_ utils/
 │
cmd/
 └─ main.go
```