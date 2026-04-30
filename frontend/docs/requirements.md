# 🧱 Frontend Application Prompt  
**Next.js (Latest) + Tailwind CSS**

## 🎯 Objective
Build a modern frontend web application using the **latest version of Next.js (App Router)** and **Tailwind CSS**, with public pages and an authenticated dashboard layout.

---

## 🛠️ Tech Stack
- **Next.js** (latest version, App Router)
- **React** (functional components)
- **Tailwind CSS**
- **TypeScript** (preferred)

---

## 🌐 Public Area (Unauthenticated Users)

Se debe visualizar una pantalla de login que permita el acceso al sistema. Este login se debe ubicar en 
app/(auth)/login/page.tsx

Que consuma el endpoint:

POST http://localhost:8090/v1/auth/login

El endpoint retorna el siguiente JSON:

{
  "access_token": "string",
  "expires_in": 86400000,
  "refresh_token": "string"
}

---

## Requisitos funcionales

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


### Public Routes
| Route | Description |
|------|------------|
| `/login` | Login page |

- **Login** must be publicly accessible.

---

## 🔐 Authentication Flow
- Login page with:
  - Email
  - Password
- Authentication call to backend to confirm the credentials.
- On successful login:
  - Redirect the user to `/dashboard`
  - Switch to an authenticated layout

---

## 📊 Authenticated Area (After Login)

### Layout
- **Header** at the top
  - Logo aligned to the left
- **Vertical sidebar menu** on the left
 - Ordenes
    - Lista
    - Estadistica
 - Logout
- **Main content area** on the right
- Sidebar must be **responsive** and **collapsible on mobile**

### Protected Routes
| Route | Description |
|------|------------|
| `/orders` | Orders list |
| `/orders/statistics` | Orders statistics |


---

## 🧩 Architectural Requirements
- Use **Next.js `layout.tsx`** to separate:
  - Public layout
  - Authenticated layout
- Reusable components:
  - Header
  - Navbar
  - Sidebar
- Clean folder structure
- Responsive design (desktop & mobile)
- Use **Tailwind utility classes only**
- No custom CSS files

---

## 📦 Deliverables
- Folder structure
- Layout files
- Reusable UI components
- Example pages
- Minimal but clean UI

---

## ✨ Optional Enhancements
- SaaS-style UI (spacing, shadows, typography)
- Dark mode support
- Route protection via middleware
- State management for authentication
