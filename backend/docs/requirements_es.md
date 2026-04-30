# Requisitos del Proyecto

Eres un Ingeniero Backend Senior especializado en Golang, diseño de APIs y aplicaciones nativas de la nube.

En nuestro equipo, gestionamos un alto volumen de transacciones diarias provenientes de diversos canales de venta (APP, Web, Tiendas Físicas). La eficiencia en el monitoreo de estos datos es crítica para la operación. El objetivo de esta API es diseñar, desarrollar y desplegar una solución que permita visualizar y analizar métricas clave a partir de un dataset de órdenes real. 

## 🧩 Requisitos Funcionales
1. Implementar autenticacion con JWT.
  - Crear una tabla para guardar usuarios (users) y contraseñas con bcrypt (password).
    - Nombre de usuario: email
    - Rol (ADMIN/USER)
    - Password
    - ID del usuario
  - Se debe validar credenciales.
  - Se debe generar:
    - Access Token (JWT) → corta duración (5–15 min)
    - Refresh Token → larga duración (días)
  - Ambos se envían como cookies seguras (HTTP-only) - Cookies HttpOnly
  - Endpoints requeridos:
    - POST /v1/users/register
      - REQUEST body:
        ```JSON
        {
          "email": "[EMAIL_ADDRESS]",
          "password": "[PASSWORD]",
          "role": "ADMIN"
        }
      ```
      - RESPONSE:
        ```JSON
        {
          "id": "1",
          "email": "[EMAIL_ADDRESS]",
          "role": "ADMIN"
        }
        ```
    - POST /v1/auth/login
      - REQUEST body:
          ```JSON
          {
            "email": "[EMAIL_ADDRESS]",
            "password": "[PASSWORD]"
          }
          ```
      - RESPONSE:
          ```JSON
          {
            "access_token": "eyJhbGciOiJIUzI1NiJ9.eyJjb21wYW55X2lkIjoiY2ZkODZiNjItZTU3ZS00ZjYwLWE1Y2ItMTRkZTlmOGM5MDlkIiwic3ViIjoiMGNhYjA1Y2YtYWQyZi00NzBmLThmOWMtODJiNDBjMGE4OGQwIiwiaWF0IjoxNzc3NTUzNzk3LCJleHAiOjE3Nzc2NDAxOTd9.5ZfADIjky9h6Ind9iX-CKY5TuF5We9yHoOkpt5k3zLo",
            "expires_in": 86400000,
            "refresh_token": "7815e5dd-e174-4d45-b197-2839da87c6cb"
          }
          ```
    - POST /v1/auth/refresh
    - REQUEST body:
        ```JSON
        {
          "refresh_token": "7815e5dd-e174-4d45-b197-2839da87c6cb"
        }
        ```
      - RESPONSE:
          ```JSON
          {
            "access_token": "eyJhbGciOiJIUzI1NiJ9.eyJjb21wYW55X2lkIjoiY2ZkODZiNjItZTU3ZS00ZjYwLWE1Y2ItMTRkZTlmOGM5MDlkIiwic3ViIjoiMGNhYjA1Y2YtYWQyZi00NzBmLThmOWMtODJiNDBjMGE4OGQwIiwiaWF0IjoxNzc3NTUzNzk3LCJleHAiOjE3Nzc2NDAxOTd9.5ZfADIjky9h6Ind9iX-CKY5TuF5We9yHoOkpt5k3zLo",
            "expires_in": 86400000,
            "refresh_token": "7815e5dd-e174-4d45-b197-2839da87c6cb"
          }
          ```
    - POST /v1/auth/logout
      - REQUEST body:
        ```JSON
        {
          "refresh_token": "7815e5dd-e174-4d45-b197-2839da87c6cb"
        }
        ```
      - RESPONSE:
        ```JSON
        {
          "message": "Logout exitoso"
        }
        ```
    - GET  /v1/auth/me
      - RESPONSE:
        ```JSON
        {
          "id": "1",
          "email": "[EMAIL_ADDRESS]",
          "role": "ADMIN"
        }
        ```

2. La api debe manejar una tabla de pedidos (orders) con los siguienes campos: 
  - id
  - canal
  - cantidad
  - company
  - cp
  - createdAt
  - daysToDelivery
  - error
  - errorMessage
  - fechaCompra
  - fechaEstimada
  - fulfillmentType
  - isFlash
  - isMarketplace
  - noPedido
  - plan
  - productType
  - sku
  - storeSelected
  - tipoPago
  - edd1
  - edd2
- Endpoints requeridos: 
  - GET /v1/orders: Debe devolver la lista de pedidos con paginación obligatoria y soporte para filtros por canal, compañía, fulfillmentType (tipo de entrega) y productType. 
  - GET /v1/stats: Debe devolver un resumen estadístico que incluya el total de órdenes, desglose por canal, desglose por tipo de entrega y por tipo de producto y el porcentaje de órdenes con err
  - Estos endpoints deben tener seguridad implementada con JWT y deben estar protegidos validando que el token sea valido.

## 🗄️ Base de Datos
- Usar PostgreSQL como base de datos.
- Generar un archivo con las query necesarias para inicializar la base de datos, se debe tomar el archivo ../docs/orders.csv y crear los inserts correspondientes.

## ⚙️ Requisitos Técnicos
- Lenguaje: Golang
- Utilizar un enfoque de arquitectura limpia (clean architecture: handlers, servicios, repositorios)
- Utilizar inyección de dependencias
- Utilizar configuración basada en variables de entorno
- Implementar manejo de errores y logging (registro de eventos) adecuados
- Seguir las convenciones de API RESTful
- Utilizar JSON para solicitudes/respuestas
- Validar las entradas
- Utilizar DTOs (Data Transfer Objects) cuando sea apropiado

## 📦 Librerías
- Enrutador (Router): chi
- Base de datos: postgresql
  - usar sqlc para generar consultas
  - nombre de base de datos : orders_db
  - puerto: 5432
  - usuario: root
  - contraseña: [PASSWORD]
- Configuración: viper
- Logging: zap

## 🔐 Mejores Prácticas
- Aplicar principios SOLID
- Estructurar el proyecto en capas (dominio, aplicación, infraestructura)
- Evitar el acoplamiento fuerte
- Escribir código limpio, mantenible y testeable
- Incluir pruebas unitarias básicas
- Validar campos requeridos y tipo de datos en la creación de usuarios

## 🧪 Extras (Opcionales)
- Dockerfile para la API
- Endpoint de verificación de salud (Health check)
- Paginación para el listado de tareas


## 📘 Requisito de Bitácora (MUY IMPORTANTE)
En cada paso del proceso de desarrollo, DEBES generar una "bitácora" explicando:

1. Qué estás haciendo
2. Por qué lo estás haciendo
3. Qué decisiones estás tomando
4. Posibles alternativas y compensaciones (trade-offs)

La bitácora debe escribirse paso a paso a medida que el proyecto evoluciona, desde el diseño inicial hasta la implementación final. Nombre del archivo: `bitacora.md` en la carpeta `./docs`.

## 📂 Estructura de Salida Esperada
1. Explicación de la arquitectura de alto nivel
2. Estructura de carpetas del proyecto
3. Esquema de base de datos + migración
4. Implementación paso a paso (con bitácora incluida)
5. Código funcional completo
6. Instrucciones para ejecutar localmente (con Docker si se incluye)
7. Ejemplos de solicitudes a la API (curl o Postman)

## 🚀 Requisitos Avanzados
- Implementar apagado controlado (graceful shutdown)
- Utilizar propagación de contexto (context propagation)
- Agregar logging estructurado con rastreo de solicitudes (request_id)
- Diseñar el sistema para que sea fácilmente extensible (por ejemplo, agregando usuarios en el futuro)
- Explicar cómo esta API podría escalar en una arquitectura de microservicios
- Explicar qué hacer en PostgreSQL para conectar la aplicación con un emulador en local, generando `postgresql-readme.md`
- Generar archivo Readme con el titulo "README.md" y que explique como correr la aplicacion y como utilizar los endpoints de la API. Debe incluir ejemplos de solicitudes a la API (curl o Postman), en este archivo se deben incluir las indicaciones para generar una migracion en sqlc, por ejemplo cuales son los pasos para agregar un campo nuevo a un tabla existente.

## 🎯 Objetivo
Entregar una API en Golang limpia y lista para producción que demuestre sólidas prácticas de ingeniería backend y un razonamiento claro a través de la bitácora.
