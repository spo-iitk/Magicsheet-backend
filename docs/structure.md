# Backend Project Structure

This document outlines the planned directory and file structure for the backend application. We are following a modular, domain-driven design pattern where each feature module (e.g., `auth`, `users`, `projects`) is self-contained with its own handlers, services, repositories, models, and routes.

```text
backend/
├── .air.toml
├── .env
├── .env.example
├── .gitignore
├── cmd
│   └── server
│       └── main.go
├── docker-compose.yml
├── docs
│   ├── schema.dbml
│   ├── schema.sql
│   └── structure.md
├── go.mod
├── go.sum
├── internal
│   ├── auth
│   │   ├── dto.go
│   │   ├── errors.go
│   │   ├── handler.go
│   │   ├── jtw.go
│   │   ├── middleware.go
│   │   ├── repository.go
│   │   ├── routes.go
│   │   └── service.go
│   ├── database
│   │   ├── db.main.go
│   │   └── db.schema.go
│   ├── magicsheet
│   │   ├── dto.go
│   │   ├── error.go
│   │   ├── handler.go
│   │   ├── mapper.go
│   │   ├── repository.go
│   │   ├── route.go
│   │   └── service.go
│   ├── middleware
│   │   ├── cors.go
│   │   ├── proforma_access.go
│   │   └── rbac.go
│   ├── rc
│   │   ├── assign.go
│   │   ├── dto.go
│   │   ├── handler.go
│   │   ├── repository.go
│   │   ├── routes.go
│   │   └── service.go
│   └── sync
│       ├── handler.go
│       ├── helper.go
│       ├── program_mapping.go
│       ├── ras_models.go
│       ├── ras_repository.go
│       ├── repository.go
│       ├── routes.go
│       └── service.go
├── migrations
└── scripts
```
