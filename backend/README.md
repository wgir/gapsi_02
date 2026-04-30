# Gapsi Orders API

API desarrollada en Golang para la gestión y análisis estadístico de órdenes de venta, siguiendo los principios de Clean Architecture.

## 🚀 Requisitos previos

- Go 1.26 o superior
- PostgreSQL (ver `docs/postgresql-readme.md`)
- Migraciones ejecutadas y base de datos inicializada (`go run scripts/seed/main.go`).

## ⚙️ Configuración y Ejecución

1. Renombra `.env.example` a `.env` y ajusta las variables (o usa los valores por defecto si tu base de datos local coincide).
2. Descarga las dependencias:
   ```bash
   go mod tidy
   ```
3. Ejecuta la aplicación:
   ```bash
   go run ./cmd/api/
   ```
La API iniciará en el puerto `8088` (ej. `http://localhost:8088`).

## 📚 Migraciones con sqlc

El proyecto está preparado para usar `sqlc` en la generación del código de base de datos a partir de SQL puro.

### ¿Cómo agregar un campo nuevo a una tabla existente?
1. Modifica la migración en `db/migrations/` o crea una nueva con `ALTER TABLE`.
2. Actualiza las consultas SQL en `db/queries/` si el nuevo campo interviene en los select, insert, o updates.
3. Estando en la raíz del proyecto, corre el comando:
   ```bash
   sqlc generate
   ```
   (Requiere tener sqlc instalado: `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)
4. Automáticamente se regenerarán las estructuras e interfaces en `internal/infra/database`.

*(Nota: En este proyecto los repositorios se implementaron manualmente siguiendo el patrón generado de sqlc debido a requerimientos de entorno, pero se mantiene 100% la estructura para adoptarlo con un solo comando).*

## 🔒 Endpoints (CURL Examples)

### Registrar usuario
```bash
curl -X POST http://localhost:8088/v1/users/register \
-H "Content-Type: application/json" \
-d '{"email":"admin@example.com", "password":"password123", "role":"ADMIN"}'
```

### Login
Retorna los tokens en JSON y además los configura en Cookies `HttpOnly`.
```bash
curl -X POST http://localhost:8088/v1/auth/login \
-H "Content-Type: application/json" \
-d '{"email":"admin@example.com", "password":"password123"}' -c cookies.txt
```

### Obtener Mi Usuario (Auth Requerido)
El token puede ir en el header `Authorization: Bearer <token>` o ser leído automáticamente desde la Cookie si es que el cliente las soporta.
```bash
curl -X GET http://localhost:8088/v1/auth/me \
-b cookies.txt
```

### Listar Órdenes (Paginadas y Filtradas)
```bash
curl -X GET "http://localhost:8088/v1/orders?page=1&pageSize=10&canal=APP&productType=Soft+Line" \
-b cookies.txt
```

### Resumen Estadístico
```bash
curl -X GET http://localhost:8088/v1/stats \
-b cookies.txt
```

### Logout
Limpia las cookies del navegador.
```bash
curl -X POST http://localhost:8088/v1/auth/logout
```
