# Configurar PostgreSQL en Local con Docker

Para correr la API necesitas una base de datos PostgreSQL. Si no tienes una instalada en tu máquina, puedes levantar un contenedor de Docker fácilmente simulando el entorno requerido por la aplicación.

## 1. Levantar PostgreSQL con Docker

Abre una terminal y ejecuta el siguiente comando:

```bash
docker run --name gapsi-postgres \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=orders_db \
  -p 5432:5432 \
  -d postgres:15-alpine
```

Esto descargará la imagen de PostgreSQL (si no la tienes) y levantará un servidor escuchando en el puerto local `5432` con las credenciales requeridas.

## 2. Verificar la conexión

Puedes verificar que el contenedor está corriendo con:

```bash
docker ps
```

Deberías ver `gapsi-postgres` en la lista. Si necesitas conectarte a la base de datos manualmente usando una herramienta como DBeaver o `psql`, usa estos datos:

- **Host**: localhost
- **Puerto**: 5432
- **Usuario**: root
- **Contraseña**: postgres
- **Base de Datos**: orders_db

## 3. Inicializar el Esquema de Base de Datos

El proyecto usa `golang-migrate` (u otro ejecutor de scripts SQL) para las tablas. Las consultas de creación (tablas `users` y `orders`) se encuentran en `db/migrations/000001_init_schema.up.sql`. Puedes ejecutar ese archivo directamente en la consola SQL de tu base de datos para crear la estructura inicial.

## 4. Cargar la data inicial desde el CSV (Seeding)

La API cuenta con un script de Go para popular la tabla de `orders` a partir de `docs/orders_db.csv`. Estando en la raíz del proyecto (`backend`), ejecuta:

```bash
go run scripts/seed/main.go
```

Esto conectará a la base de datos `orders_db` y llenará la tabla de órdenes con los registros.
