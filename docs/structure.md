# Backend Project Structure

This document outlines the planned directory and file structure for the backend application. We are following a modular, domain-driven design pattern where each feature module (e.g., `auth`, `users`, `projects`) is self-contained with its own handlers, services, repositories, models, and routes.

```text
backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── auth/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   ├── middleware.go
│   │   ├── routes.go
│   │   └── dto.go
│   │
│   ├── database/
│   ├── config/
│   ├── utils/
│   └── common/
│
├── migrations/
├── docs/
└── go.mod
```
