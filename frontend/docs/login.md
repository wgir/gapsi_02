# Prompt – Implementación Login en Next.js 16.1.6 (App Router)

Actúa como un desarrollador senior especializado en **Next.js 16.1.6 (App Router)**, autenticación JWT y buenas prácticas de seguridad.

Necesito implementar una página:

app/(auth)/login/page.tsx

Que consuma el endpoint:

POST http://localhost:8080/v1/auth/login

El endpoint retorna el siguiente JSON:

{
  "access_token": "string",
  "expires_in": 86400000,
  "refresh_token": "string"
}

---

## Requisitos funcionales

1. Usar la page app/(auth)/login/page.tsx para la integracion

2. En el hook app/app/useAuth.ts enviar los datos al endpoint usando `fetch`, quitendo el mock existente.

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

---

## Requisitos técnicos

- Usar App Router (NO Pages Router).
- Usar Server Actions o Route Handlers.
- No exponer tokens en el cliente.
- Usar TypeScript estricto.
- Separar responsabilidades (UI, lógica, infraestructura).
- Código listo para producción.
- Implementar validación básica del formulario.

---

## Entregables esperados

1. Código completo de:
   - app/(auth)/login/page.tsx
   - app/app/useAuth.ts arreglado para que funcione con el login

2. Middleware para proteger rutas privadas (`proxy.ts`).

3. Explicación breve de:
   - Estrategia de almacenamiento del token.
   - Flujo de autenticación.
   - Manejo de refresh token automático.

---

## Extras (si aplica)

- Implementar patrón enterprise-ready.
- Manejar refresh token silencioso.
- Buenas prácticas de seguridad (SameSite, Secure, httpOnly).
- Estructura recomendada del proyecto.

Genera código limpio, profesional y bien estructurado.


--------------------------------------------------------------------------------