import axios from 'axios';
import Cookies from 'js-cookie';

const API_BASE_URL = '/api';

export const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add a request interceptor to add the access token to headers
api.interceptors.request.use(
  (config) => {
    const token = Cookies.get('access_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Add a response interceptor to handle token expiration
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      if (typeof window !== 'undefined') {
        // Prevent redirect loop if already on login page or already logging out
        if (!window.location.pathname.includes('/login') && !window.location.pathname.includes('/logout')) {
          // Redirect to logout route to clear httpOnly cookies server-side
          window.location.href = '/api/auth/logout';
        }
      }
    }

    return Promise.reject(error);
  }
);

export default api;
