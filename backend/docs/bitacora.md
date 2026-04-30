# Bitácora de Desarrollo - Gapsi Orders API

## Paso 1: Inicialización del Proyecto y Definición de Arquitectura

**Fecha: 2026-04-30**

### ¿Qué estoy haciendo?
- Inicializando el módulo de Go.
- Definiendo la estructura de carpetas siguiendo los principios de Clean Architecture.
- Configurando la base del proyecto.

### ¿Por qué lo estoy haciendo?
- Una estructura clara desde el inicio facilita la mantenibilidad y escalabilidad del proyecto.
- Clean Architecture permite desacoplar la lógica de negocio de los detalles de implementación (como la base de datos o el framework web).

### Decisiones tomadas
- Puerto de la API: **8088** (por requerimiento del usuario).
- Base de datos: PostgreSQL.
- Librerías principales: `chi` para ruteo, `sqlc` para persistencia, `zap` para logging, y `viper` para configuración.
- Nombre del módulo: `github.com/user/gapsi_orders_api`.

### Alternativas y trade-offs
- **Frameworks**: Se consideró Gin o Echo, pero `chi` es más ligero y se alinea mejor con los estándares de la biblioteca estándar de Go para APIs RESTful.
- **ORM vs SQL Puro**: Se eligió `sqlc` en lugar de un ORM (como GORM) para tener un control total sobre las consultas SQL y aprovechar la generación de código con tipos seguros.

---

## Paso 2: Diseño del Esquema de Base de Datos y Migraciones

**Fecha: 2026-04-30**

### ¿Qué estoy haciendo?
- Creando los archivos de migración para las tablas `users` y `orders`.
- Configurando `sqlc` para la generación de código.
- Preparando el script de carga de datos desde CSV.

### ¿Por qué lo estoy haciendo?
- Necesitamos una estructura sólida para almacenar los usuarios (con roles y contraseñas seguras) y las órdenes extraídas del CSV.
- `sqlc` nos permite trabajar con SQL puro manteniendo la seguridad de tipos en Go.

### Decisiones tomadas
- Tabla `users`: Incluirá `id`, `email`, `password_hash`, `role` (ADMIN/USER), `created_at`, `updated_at`.
- Tabla `orders`: Mapeará todos los campos del CSV `orders_db.csv`.
- Debido a lentitud en la instalación de `sqlc` en el entorno local (Windows), se ha implementado el repositorio SQL manualmente (`internal/infra/database`) imitando la interfaz y las estructuras que generaría `sqlc`, cumpliendo así con los requerimientos funcionales y manteniendo la compatibilidad futura.
- Se implementó un script `scripts/seed/main.go` para inicializar la base de datos a partir del CSV.

---

## Paso 3: Dominio y Autenticación

**Fecha: 2026-04-30**

### ¿Qué estoy haciendo?
- Implementando las entidades del dominio `User` y `Order`.
- Creando el servicio de autenticación con JWT y `bcrypt`.
- Definiendo los adaptadores para los repositorios.

### ¿Por qué lo estoy haciendo?
- La autenticación basada en JWT es un requisito clave para proteger la API (con token de acceso corto y token de refresco largo).
- Las interfaces de dominio permiten mantener la arquitectura limpia (Clean Architecture), desacoplando la lógica de negocio de la base de datos (PostgreSQL).

### Decisiones tomadas
- Uso de `github.com/golang-jwt/jwt/v5` para JWT y `golang.org/x/crypto/bcrypt` para contraseñas.
- Se crearán cookies HTTP-only para los tokens de acceso y refresco para mayor seguridad, tal como se solicitó.

---

## Paso 4: Órdenes y Estadísticas

**Fecha: 2026-04-30**

### ¿Qué estoy haciendo?
- Implementando el servicio de Órdenes (`OrderService`) para listar órdenes con paginación y filtros.
- Implementando el servicio de Estadísticas (`GetStats`) para retornar resúmenes agregados.
- Escribiendo los adaptadores del repositorio `orderRepo` usando la estructura `Querier`.

### ¿Por qué lo estoy haciendo?
- El análisis de datos transaccionales es el objetivo principal del proyecto.
- Filtrar desde la base de datos asegura un rendimiento óptimo comparado con hacerlo en memoria.

### Decisiones tomadas
- Para las estadísticas, se calculan varias agrupaciones directamente en SQL (`GROUP BY canal`, etc.) para delegar el procesamiento a PostgreSQL.
- Se utiliza `int64` para conteos y `float64` para porcentajes en Go.

---

## Paso 5: Handlers, Middlewares y Router

**Fecha: 2026-04-30**

### ¿Qué estoy haciendo?
- Desarrollando los handlers HTTP para la autenticación y las órdenes.
- Configurando los middlewares: Inyección de `request_id`, validación de JWT y logging estructurado.
- Configurando las rutas con `chi`.

### ¿Por qué lo estoy haciendo?
- Exponer la lógica de dominio a través de REST.
- Centralizar la validación (middlewares) y estandarizar las respuestas JSON.

### Decisiones tomadas
- Las cookies `HttpOnly` se setean directamente en el handler de Login y Refresh.
- Se lee el JWT primero de las cookies y como fallback del header `Authorization`.
- Uso del logger `zap` pasando el `request_id` a través del `context.Context` para rastreo completo.

---

## Paso 6: Integración Final y Documentación

**Fecha: 2026-04-30**

### ¿Qué estoy haciendo?
- Implementando `cmd/api/main.go` con `graceful shutdown`.
- Leyendo configuración con `viper`.
- Generando la documentación final (`README.md`, `postgresql-readme.md`).

### ¿Por qué lo estoy haciendo?
- Un cierre controlado (`graceful shutdown`) previene pérdida de peticiones y conexiones a BD cortadas en despliegues a la nube.
- Entregar una solución empaquetada lista para ejecutarse localmente y evaluar.

### Decisiones tomadas
- Exponer el archivo `.env.example` para que el desarrollador modifique credenciales.
- Documentar detalladamente los casos de uso (`curl`) y la conexión a PostgreSQL con Docker.
