# Gapsi Orders Dashboard

Aplicación full-stack para la visualización, filtrado y gestión de órdenes de venta, construida con una arquitectura moderna de microservicios usando **Golang** en el backend y **Next.js 16** en el frontend.

## 🚀 Requisitos Previos

Asegúrate de tener instalados los siguientes componentes antes de iniciar:

- **Go** (1.26+)
- **Node.js** (v18+ recomendado)
- **Docker** y **Docker Compose** (opcional, recomendado para levantar la base de datos)
- **PostgreSQL** (si no usas Docker)

---

## 📂 Estructura del Proyecto

- `backend/`: API RESTful en Go, manejando autenticación JWT, conexión a base de datos (PostgreSQL), filtrado avanzado y paginación. Usa arquitectura limpia (Clean Architecture).
- `frontend/`: Aplicación web responsiva en Next.js App Router, con Tailwind CSS para estilos, React Query para manejo de estado asíncrono y Recharts para estadísticas visuales.

---

## ⚙️ Guía de Ejecución Local

Sigue estos pasos en orden para levantar el sistema completo de forma local.

### 1. Levantar la Base de Datos (PostgreSQL)

La forma más rápida de tener la base de datos lista es usando Docker. Si tienes Docker instalado, ejecuta el siguiente comando para levantar un contenedor de PostgreSQL:

```bash
docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=orders_db -d postgres:alpine
```

### 2. Configuración y Ejecución del Backend (Go)

Abre una terminal y dirígete a la carpeta `backend`:

```bash
cd backend
```

1. **Configurar el entorno**: Copia el archivo `.env.example` a `.env` (si aún no existe) y asegúrate de que las variables de la base de datos coincidan con las de tu contenedor:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=root
   DB_PASSWORD=secret
   DB_NAME=orders_db
   SERVER_PORT=8090
   ```

2. **Instalar dependencias y poblar base de datos**:
   ```bash
   go mod tidy
   go run scripts/seed/main.go
   ```
   *(Este último comando creará las tablas y las poblará con la data de prueba).*

3. **Iniciar la API**:
   ```bash
   go run ./cmd/api/
   ```
   *La API estará corriendo en http://localhost:8090.*

4. **Crear usuario administrador**:
   Abre una nueva terminal y crea un usuario inicial mediante cURL (o Postman):
   ```bash
   curl -X POST http://localhost:8090/v1/users/register \
   -H "Content-Type: application/json" \
   -d '{"email":"admin@example.com", "password":"password123", "role":"ADMIN"}'
   ```

### 3. Configuración y Ejecución del Frontend (Next.js)

Abre una nueva terminal y dirígete a la carpeta `frontend`:

```bash
cd frontend
```

1. **Configurar el entorno**: Asegúrate de tener o crear el archivo `.env.local` en la raíz del frontend para apuntar a la API:
   ```env
   API_URL=http://localhost:8090
   ```

2. **Instalar dependencias**:
   ```bash
   npm install
   ```

3. **Iniciar el servidor de desarrollo**:
   ```bash
   npm run dev
   ```

---

## 🌐 Uso del Sistema

Una vez que ambos servidores (Backend y Frontend) estén corriendo:

1. Abre tu navegador y dirígete a: **`http://localhost:3000`**
2. Se te redirigirá a la pantalla de inicio de sesión.
3. Ingresa con las credenciales que creaste en el paso del backend:
   - **Email:** `admin@example.com`
   - **Password:** `password123`
4. ¡Listo! Ya tienes acceso al Dashboard completo donde podrás ver los gráficos interactivos, estadísticas globales y la tabla de órdenes con sus respectivos filtros y paginación.
