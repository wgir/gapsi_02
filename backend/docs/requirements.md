You are a Senior Backend Engineer specialized in Golang, API design, and cloud-native applications.

En nuestro equipo, gestionamos un alto volumen de transacciones diarias provenientes de diversos canales de venta (APP, Web, Tiendas Físicas). La eficiencia en el monitoreo de estos datos es crítica para la operación. El objetivo de esta api es diseñar, desarrollar y desplegar una solución que permita visualizar y analizar métricas clave a partir de un dataset de órdenes real. 

## 🧩 Functional Requirements
- La api debe manejar una tabla de pedidos con los siguienes campos: 
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
  - GET /api/orders: Debe devolver la lista de pedidos con paginación obligatoria y soporte para filtros por canal, compañía, fulfillmentType (tipo de entrega) y productType. 
  - GET /api/stats: Debe devolver un resumen estadístico que incluya el total de órdenes, desglose por canal, desglose por tipo de entrega y por tipo de producto y el porcentaje de órdenes con err
  

## 🗄️ Database
- Use PostgreSQL as database

## ⚙️ Technical Requirements
- Language: Golang
- Use a clean architecture approach (handlers, services, repositories)
- Use dependency injection
- Use environment-based configuration
- Implement proper error handling and logging
- Follow RESTful API conventions
- Use JSON for request/response
- Validate inputs
- Use DTOs where appropriate

## 📦 Libraries
- Router: chi
- Database: postgresql
- Config: viper
- Logging: zap

## 🔐 Best Practices
- Apply SOLID principles
- Structure project in layers (domain, application, infrastructure)
- Avoid tight coupling
- Write clean, maintainable, and testable code
- Include basic unit tests

## 🧪 Extras (Nice to have)
- Dockerfile for the API
- Health check endpoint
- Pagination for task listing


## 📘 Bitácora Requirement (VERY IMPORTANT)
At each step of the development process, you MUST generate a "bitácora" (logbook) explaining:

1. What you are doing
2. Why you are doing it
3. What decisions you are making
4. Possible alternatives and trade-offs

The bitácora must be written step-by-step as the project evolves, from initial design to final implementation. File name: bitacora.md in ./docs folder.

## 📂 Expected Output Structure
1. High-level architecture explanation
2. Project folder structure
3. Database schema + migration
4. Step-by-step implementation (with bitácora included)
5. Complete working code
6. Instructions to run locally (with Docker if included)
7. Example API requests (curl or Postman)

## 🚀 Advanced Requirements
- Implement graceful shutdown
- Use context propagation
- Add structured logging with request tracing (request_id)
- Design the system to be easily extensible (e.g., adding users in the future)
- Explain how this API could scale in a microservices architecture
- Explain what to do in firestore to connect the application with an emulator on local, generating firestore-readme.md

## 🎯 Goal
Deliver a clean, production-ready Golang API that demonstrates strong backend engineering practices and clear reasoning through the bitácora.