# Gapsi Orders API

API desarrollada en Golang para la gestión y análisis estadístico de órdenes de venta, siguiendo los principios de Clean Architecture.

## 🚀 Requisitos previos

- Go 1.26 o superior
- PostgreSQL (ver `docs/postgresql-readme.md`)
- Migraciones ejecutadas y base de datos inicializada (`go run scripts/seed/main.go`).

## ⚙️ Configuración y Ejecución

1. En `.env` ajusta las variables de ambiente según tu entorno local.
2. Descarga las dependencias:
   ```bash
   go mod tidy
   ```
3. Inicializa la base de datos ejecutando el script (esto creará las tablas y poblará los datos del CSV):
   ```bash
   go run scripts/seed/main.go
   ```
4. Ejecuta la aplicación (la API iniciará en el puerto `8090`):
   ```bash
   go run ./cmd/api/
   ```
5. Registra tu primer usuario administrador para poder ingresar al sistema. Puedes usar Postman o ejecutar el siguiente comando cURL en otra terminal:
   ```bash
   curl -X POST http://localhost:8090/v1/users/register \
   -H "Content-Type: application/json" \
   -d '{"email":"admin@example.com", "password":"password123", "role":"ADMIN"}'
   ```

¡Con esto ya tienes el backend 100% operativo y puedes iniciar sesión en el dashboard con las credenciales que acabas de crear!

## 📚 Migraciones con sqlc

El proyecto está preparado para usar `sqlc` en la generación del código de base de datos a partir de SQL puro.

### ¿Cómo agregar un campo nuevo a una tabla existente?
1. Modifica la migración en `db/migrations/` o crea una nueva con `ALTER TABLE`.
2. Actualiza las consultas SQL en `db/queries/` si el nuevo campo interviene en los select, insert, o updates.
3. Estando en la raíz del proyecto, corre el comando:
   ```bash
   sqlc generate 
   ```
   (Opcional: `go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1 generate`)
4. Automáticamente se regenerarán las estructuras e interfaces en `internal/infra/database/sqlc`.


## 🔒 Endpoints (CURL Examples)

### Registrar usuario
```bash
curl -X POST http://localhost:8090/v1/users/register \
-H "Content-Type: application/json" \
-d '{"email":"admin@example.com", "password":"password123", "role":"ADMIN"}'
```

### Login
Retorna los tokens en JSON y además los configura en Cookies `HttpOnly`.
```bash
curl -X POST http://localhost:8090/v1/auth/login \
-H "Content-Type: application/json" \
-d '{"email":"admin@example.com", "password":"password123"}' -c cookies.txt
```

### Obtener Mi Usuario (Auth Requerido)
El token puede ir en el header `Authorization: Bearer <token>` o ser leído automáticamente desde la Cookie si es que el cliente las soporta.
```bash
curl -X GET http://localhost:8090/v1/auth/me \
-b cookies.txt
```

### Listar Órdenes (Paginadas y Filtradas)
```bash
curl -X POST http://localhost:8090/v1/orders \
-H "Content-Type: application/json" \
-b cookies.txt \
-d '{
  "page": 1,
  "page_size": 10,
  "canal": "APP",
  "product_type": "Soft Line"
}'
```

### Resumen Estadístico
```bash
curl -X GET http://localhost:8090/v1/stats \
-b cookies.txt
```

### Logout
Limpia las cookies del navegador.
```bash
curl -X POST http://localhost:8090/v1/auth/logout
```

## 🐳 Uso con Docker

El proyecto cuenta con un `Dockerfile` optimizado en múltiples etapas (multi-stage build) para generar una imagen ligera y lista para producción.

### 1. Construir la imagen
Posiciónate en la carpeta `backend` y ejecuta:
```bash
docker build -t gapsi-orders-api .
```

### 2. Ejecutar el contenedor
Puedes correr el contenedor vinculándolo directamente a tu contenedor de base de datos (por ejemplo, si se llama `postgres-container`):

```bash
 docker run -p 8090:8090 --link postgres-container:db -e SERVER_PORT=8090 -e DB_HOST=db -e DB_PORT=5432 -e DB_USER=root -e DB_PASSWORD=secret -e DB_NAME=orders_db -e DB_SSLMODE=disable -e JWT_SECRET=super_secret_key_change_me_in_production -e JWT_ACCESS_TTL_MINUTES=75 -e JWT_REFRESH_TTL_DAYS=7 gapsi-orders-api
```

*(La API estará disponible en `http://localhost:8090`)*
