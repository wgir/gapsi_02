import axios from 'axios';

/**
 * Axios instance that targets the Next.js BFF routes (/api/*).
 *
 * NOTE: Authentication tokens are stored as httpOnly cookies and are
 * therefore NOT accessible from JavaScript.  The cookies are forwarded
 * automatically by the browser for same-origin requests, so the BFF
 * route handlers read them server-side via next/headers.
 *
 * We do NOT add an Authorization header here — it would always be empty.
 * Token forwarding to the real backend happens inside each /app/api/* route.
 */
const api = axios.create({
  baseURL: '/api',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Handle global 401s: redirect to logout so the server clears httpOnly cookies
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (
      error.response?.status === 401 &&
      typeof window !== 'undefined' &&
      !window.location.pathname.includes('/login')
    ) {
      window.location.href = '/api/auth/logout';
    }
    return Promise.reject(error);
  },
);

export { api };
export default api;
