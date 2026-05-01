'use client';

import { useAuth } from '@/hooks/useAuth';
import { LogOut, LayoutDashboard } from 'lucide-react';
import Link from 'next/link';

/**
 * Sidebar — Client Component only because it needs useAuth (logout action).
 * Kept isolated so the parent DashboardLayout stays a Server Component.
 */
export default function SidebarClient() {
  const { logout } = useAuth();

  return (
    <aside className="w-64 bg-white border-r border-gray-200 hidden md:flex flex-col">
      <div className="p-6 border-b border-gray-100">
        <h1 className="text-xl font-bold text-blue-600 flex items-center gap-2">
          <LayoutDashboard className="h-6 w-6" />
          Gapsi Admin
        </h1>
      </div>

      <nav className="flex-1 p-4 space-y-1">
        <Link
          href="/dashboard"
          className="flex items-center gap-3 px-4 py-3 text-sm font-medium text-blue-600 bg-blue-50 rounded-lg"
        >
          <LayoutDashboard className="h-5 w-5" />
          Dashboard
        </Link>

        <button
          onClick={logout}
          className="flex w-full items-center gap-3 px-4 py-3 text-sm font-medium text-red-600 hover:bg-red-50 rounded-lg transition-colors"
        >
          <LogOut className="h-5 w-5" />
          Cerrar sesión
        </button>
      </nav>
    </aside>
  );
}
