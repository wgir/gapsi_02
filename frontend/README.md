# Gapsi Orders Dashboard - Frontend

Este es el frontend de la aplicación Gapsi Orders, construido con **Next.js 16 (App Router)**, **TypeScript**, y **Tailwind CSS**.

## 🛠️ Tecnologías Principales

- **Framework**: [Next.js 16](https://nextjs.org/)
- **Estilos**: [Tailwind CSS v4](https://tailwindcss.com/)
- **Manejo de Estado Global**: [Zustand](https://zustand-demo.pmnd.rs/)
- **Obtención de Datos (Data Fetching)**: [TanStack React Query v5](https://tanstack.com/query/latest)
- **Gráficos Estadísticos**: [Recharts](https://recharts.org/)
- **Iconos**: [Lucide React](https://lucide.dev/)

---

## ⚙️ Configuración del Entorno

Antes de iniciar la aplicación, necesitas configurar las variables de entorno para que el frontend sepa dónde comunicarse con el backend (API).

1. En la raíz de esta carpeta (`frontend/`), crea un archivo llamado `.env.local`.
2. Añade la variable `API_URL` apuntando a la dirección de tu servidor backend. Por defecto:
   ```env
   API_URL=http://localhost:8090
   ```

---

## 🚀 Instalación y Ejecución

Asegúrate de tener **Node.js** (versión 18 o superior) instalado.

### 1. Instalar dependencias
Abre una terminal en esta carpeta (`frontend`) y ejecuta:
```bash
npm install
```

### 2. Iniciar el servidor de desarrollo
Para correr la aplicación en modo desarrollo con recarga en caliente (Hot Reload):
```bash
npm run dev
```
La aplicación estará disponible en [http://localhost:3000](http://localhost:3000).

---

## 📦 Construcción para Producción

Si deseas generar una versión optimizada y lista para producción:

1. **Construir el proyecto**:
   ```bash
   npm run build
   ```
2. **Iniciar el servidor de producción**:
   ```bash
   npm run start
   ```

*(Nota: Para producción, también debes configurar tus variables de entorno correspondientes).*
