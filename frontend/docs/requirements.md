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

### Layout
- **Header** displayed at the top of the page
  - Logo aligned to the **left**
  - Navigation bar aligned to the **right**

### Navigation Links
- Login

### Public Routes
| Route | Description |
|------|------------|
| `/` | Login page |

- **Login** must be publicly accessible.

---

## 🔐 Authentication Flow
- Login page with:
  - Email
  - Password
- Authentication can be **mocked** (no backend required).
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
