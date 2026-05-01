'use client';

import { useAuth } from '@/hooks/useAuth';
import { LogOut } from 'lucide-react';

/** Mobile-only logout button shown in the top header. */
export default function LogoutButtonClient() {
  const { logout } = useAuth();

  return (
    <button
      onClick={logout}
      className="md:hidden p-2 text-gray-500 hover:text-red-600 transition-colors"
      aria-label="Cerrar sesión"
    >
      <LogOut className="h-5 w-5" />
    </button>
  );
}
