import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { AuthResponse, LoginError } from '@/types/auth';

export const useAuth = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<LoginError | null>(null);
  const router = useRouter();

  const login = async (email: string, password: string) => {
    setIsLoading(true);
    setError(null);

    try {
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        // Server returned 4xx/5xx — parse the error message from the body
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || `Error ${response.status}: no se pudo iniciar sesión`);
      }

      // Login succeeded — the server set the httpOnly cookies.
      // Redirect unconditionally; do NOT rely on a `success` flag in the body.
      router.push('/dashboard');
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Error desconocido al iniciar sesión';
      setError({ message });
    } finally {
      setIsLoading(false);
    }
  };

  const logout = async () => {
    try {
      await fetch('/api/auth/logout', { method: 'POST' });
    } catch (err) {
      // Best-effort: even if the network request fails, redirect to login
      console.error('Logout request failed:', err);
    } finally {
      router.push('/login');
    }
  };

  return { login, logout, isLoading, error };
};
