Desarrollo de Frontend - Dashboard de Órdenes (React 19 + Next.js 16)
🎯 Objetivo

Construir un dashboard web profesional utilizando React 19 + Next.js 16 que permita visualizar, filtrar y analizar órdenes de venta en tiempo real, consumiendo una API backend.

🧱 Stack Tecnológico
Framework: Next.js (App Router recomendado)
Lenguaje: TypeScript
Estilos: Tailwind CSS
Manejo de datos: React Query (TanStack Query)
Gráficos: Recharts o Chart.js
Manejo de estado global: Zustand o Context API
Cliente HTTP: fetch o Axios
🔐 Autenticación

Implementar un módulo de autenticación basado en JWT:

Requisitos:
Página /login
Formulario: email + password
Manejo de sesión con cookies HTTP-only
Redirección a /dashboard tras login exitoso
Protección de rutas (middleware o lógica en layout)

endpoint:

POST http://localhost:8090/v1/auth/login

El endpoint retorna el siguiente JSON:

{
  "access_token": "string",
  "expires_in": 86400000,
  "refresh_token": "string"
}

---

## Requisitos funcionales del login

1. Usar la page app/(auth)/login/page.tsx para la integracion

2. En el hook app/app/useAuth.ts enviar los datos al endpoint usando `fetch`.

3. Manejar estados:
   - loading
   - error
   - success

4. Si el login es exitoso:
   - Guardar `access_token` en una cookie segura (`httpOnly` usando Route Handler).
   - Guardar `refresh_token`.
   - Configurar expiración usando `expires_in`.
   - Redirigir al usuario a `/dashboard`.

5. Manejar correctamente errores HTTP:
   - 400
   - 401
   - 500


🔑 Comportamiento inicial
La pantalla inicial de la aplicación debe ser /login
Si el usuario NO está autenticado, siempre debe ser redirigido a /login
Si el usuario YA está autenticado, debe ser redirigido automáticamente a /dashboard

🧭 Estructura de Rutas
/app
  /login
  /dashboard
    page.tsx
    layout.tsx
/components
/hooks
/services
/types

🖥️ Dashboard (Vista Principal)

Construir una single-page dashboard en /dashboard con los siguientes módulos:

1. 🔍 Filtros (Globales)
canal
compañía
fulfillmentType
productType

📌 Deben afectar:

tabla
gráficos
estadísticas
2. 📊 Estadísticas (KPIs)

Consumir:

GET /api/stats

Mostrar:

Total de órdenes
% de órdenes con error
Desglose por canal
Desglose por tipo de entrega
Desglose por tipo de producto
3. 📈 Gráficos

Mínimo 2 gráficos:

Órdenes por canal
Órdenes por tipo de entrega o producto

📌 Deben actualizarse en tiempo real con los filtros

el endpoint es:
GET http://localhost:8090/v1/orders/stats

ejemplo 
```
'http://localhost:8090/v1/orders/stats' \
  --header 'Content-Type: application/json' \
  --header 'Cookie: access_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTY2MjIwOWEtODg3NC00Y2IxLWI0YzQtYWZjNTA0OGMzODFjIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInJvbGUiOiJBRE1JTiIsInR5cGUiOiJhY2Nlc3MiLCJleHAiOjE3Nzc1Nzk3OTIsImlhdCI6MTc3NzU3ODg5Mn0.ly8LXA4dLhZYxDIvOPUKjjgKF4qkZfs-xHdTU5uf8_Y; refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTY2MjIwOWEtODg3NC00Y2IxLWI0YzQtYWZjNTA0OGMzODFjIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInJvbGUiOiJBRE1JTiIsInR5cGUiOiJyZWZyZXNoIiwiZXhwIjoxNzc4MTgzNjkyLCJpYXQiOjE3Nzc1Nzg4OTJ9._QErTQlt3qewRFZGjC6SMvmV9-B9keyy5lCwt-mG2cs'`
````

Response :

```json
{
    "total_orders": 59848,
    "breakdown_by_canal": {
        "": 1,
        "APP": 6069,
        "APV": 1,
        "APV AND": 24008,
        "MOBILE": 1237,
        "WEB": 28532
    },
    "breakdown_by_fulfillment": {
        "": 22439,
        "Fulfillment_Type_Liverpool": 30360,
        "Liverpool_CNC_PICK_PACK": 7049
    },
    "breakdown_by_product_type": {
        "": 22439,
        "BT": 345,
        "Big Ticket": 9431,
        "SL": 1226,
        "Soft Line": 26407
    },
    "percentage_with_errors": 12.018780911642828
}
```

4. 📋 Tabla de Órdenes

Consumir:

GET /api/orders

Requisitos:

Paginación obligatoria
Columnas:
noPedido
canal
sku
fechaEstimada
tipoEntrega
tipoProducto
cantidad
fechaCompra

Endpoint: 
GET http://localhost:8090/v1/orders?page=1&page_size=5&canal=APP&company=1&fulfillment_type=FFW&&product_type=1

El endpoint retorna el siguiente JSON:
{
  "code": "ok",
  "message": "ok",
  "data": {
    "total_rows": 24840,
    "total_pages": 4968,
    "current_page": 1,
    "page_size": 5,
    "data": [
      {
        "id": 34194826,
        "code": "108-2696996-1831409",
        "canal": "APP",
        "company": "BODEGA-01",
        "fulfillmentType": "FFW",
        "productType": "FOOD",
        "salesOrderId": 4053789,
        "orderStatus": {
          "code": "PENDIENTE_PAGO",
          "description": "Pendiente de pago",
          "id": 1202,
          "can_cancel": true
        },
        "dateCreated": "2026-04-23T14:46:02.000-05:00",
        "paymentId": "0172F25A-E7DF-7C79-B5A8-74942DF2F5F4",
        "buyerUserId": 42521609,
        " buyerUserEmail": "[EMAIL_ADDRESS]",
        "paymentStatus": {
          "code": "CONFIRMADO",
          "description": "Proceso completado",
          "id": 3,
          "can_cancel": false
        },
        "totalPayment": "0.00",
        "orderNotes": null,
        "sellerUserEmail": "[EMAIL_ADDRESS]",
        "sellerCompanyCode": "311-7810245-101",
        "shipmentDate": "2026-04-28T00:00:00.000-05:00",
        "deliveryDate": null,
        "giftCardNumber": null,
        "orderAddress": {
          "id": 48924805,
          "apartment": null,
          "streetAddress": "",
          "addressLine2": null,
          "postalCode": null,
          "addressName": "TEST",
          "latitude": null,
          "longitude": null,
          "phone": "",
          "floor": null,
          "district": null,
          "city": "",
          "state": "",
          "country": "CO",
          "neighborhood": null,
          "reference": null
        },
        "customerDetails": {
          "id": 2683997,
          "userId": 42521609,
          "documentType": {
            "id": 2,
            "code": "CED",
            "description": "Cédula"
          },
          "documentNumber": "92316544",
          "name": "TEST",
          "email": "[EMAIL_ADDRESS]"
        },
        "paymentMethod": {
          "id": 10,
          "code": "null",
          "description": " null",
          "paymentMethodType": {
            "id": 1,
            "code": "CREDIT_CARD",
            "description": " tarjeta de crédito"
          },
          "isOnline": true,
          "canCancel": true
        },
        "deliveryMethod": {
          "id": 3,
          "code": "DELIVERY_AT_HOME",
          "description": "Entrega en el domicilio",
          "canCancel": false
        },
        "totalPaymentsWithoutTaxes": "0.00",
        "totalPaymentsWithTaxes": "0.00",
        "orderStatusHistory": [
          {
            "id": 265351462,
            "status": {
              "code": "PENDIENTE_PAGO",
              "description": "Pendiente de pago",
              "id": 1202,
              "can_cancel": true
            },
            "isCurrent": true,
            "dateCreated": "2026-04-23T14:46:02.000-05:00",
            "userEmail": "[EMAIL_ADDRESS]"
          },
          {
            "id": 265351464,
            "status": {
              "code": "PENDIENTE_ENVIO",
              "description": "Pendiente de envío",
              "id": 10,
              "can_cancel": true
            },
            "isCurrent": false,
            "dateCreated": "2026-04-23T14:46:02.000-05:00",
            "userEmail": "[EMAIL_ADDRESS]"
          },
          {
            "id": 265351465,
            "status": {
              "code": "ENTREGADO",
              "description": "Entregado",
              "id": 15,
              "can_cancel": false
            },
            "isCurrent": false,
            "dateCreated": "2026-04-23T14:46:02.000-05:00",
            "userEmail": "[EMAIL_ADDRESS]"
          }
        ],
        "orderItems": [
          {
            "id": 223735440,
            "orderCode": "108-2696996-1831409",
            "orderId": 34194826,
            "sku": "99910695",
            "name": "Bocadillo tarrito x 10 unds",
            "quantity": 2,
            "unitPrice": "0.00",
            "unitPriceWithTaxes": "0.00",
            "lineTotal": "0.00",
            "lineTotalWithTaxes": "0.00",
            "isCombo": true,
            "comboItems": [
              {
                "id": 447470879,
                "name": "PAÑUELO Facial Huggies 10 Unidades",
                "sku": "7702425776876",
                "quantity": 2,
                "isComboItem": true,
                "unitPrice": "0.00",
                "unitPriceWithTaxes": "0.00",
                "lineTotal": "0.00",
                "lineTotalWithTaxes": "0.00"
              }
            ]
          }
        ]
      }
    ],
    "page": 1,
    "page_size": 5,
    "total": 6069
  }
}

🔄 Manejo de Estado
Centralizar filtros en estado global
Usar React Query para:
cache
refetch automático
sincronización con filtros
🧩 Arquitectura sugerida
/components
  /filters
  /charts
  /table
  /ui
/hooks
  useOrders.ts
  useStats.ts
  useFilters.ts
/services
  api.ts
/types
  order.ts
  stats.ts
🎨 UI/UX
Diseño limpio, moderno y responsivo
Mobile-first
Uso consistente de spacing y tipografía
Skeleton loaders / loading states
Manejo de errores visual
⚡ Requisitos de Calidad
Código modular y escalable
Uso de TypeScript estricto
Separación clara de responsabilidades
Evitar lógica duplicada
Manejo correcto de loading/error states
🚀 Bonus (Opcional pero recomendado)
Dark mode
Persistencia de filtros en URL (query params)
Debounce en filtros
Exportación de datos (CSV)
Testing básico (React Testing Library)
🏁 Resultado Esperado

Un dashboard interactivo donde:

El usuario puede autenticarse
La app inicia siempre en login si no hay sesión
Visualiza métricas clave
Filtra datos en tiempo real
Analiza órdenes en una sola vista integrada
⚠️ Consideraciones Clave
NO separar tabla y gráficos en diferentes páginas
TODO debe vivir en /dashboard
Filtros deben ser globales y sincronizados
La autenticación debe controlar el acceso a toda la app