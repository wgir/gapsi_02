'use client';

import { useAuth } from '@/hooks/useAuth';
import { LogOut, LayoutDashboard, ShoppingCart, BarChart3, User } from 'lucide-react';
import Link from 'next/link';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { logout } = useAuth();

  return (
    <div className="flex min-h-screen bg-gray-50">
      {/* Sidebar */}
      <aside className="w-64 bg-white border-r border-gray-200 hidden md:block">
        <div className="h-full flex flex-col">
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

          <div className="p-4 border-t border-gray-100">

          </div>
        </div>
      </aside>

      {/* Main Content */}
      <div className="flex-1 flex flex-col">
        {/* Top Header */}
        <header className="h-16 bg-white border-b border-gray-200 flex items-center justify-between px-8">
          <div className="md:hidden">
            <h1 className="text-xl font-bold text-blue-600">Gapsi</h1>
          </div>
          <div className="flex-1 hidden md:block">
            {/* Search or breadcrumbs could go here */}
          </div>
          <div className="flex items-center gap-4">
            <div className="flex items-center gap-2 px-3 py-1.5 bg-gray-50 rounded-full border border-gray-100">
              <div className="h-8 w-8 rounded-full bg-blue-100 flex items-center justify-center">
                <User className="h-5 w-5 text-blue-600" />
              </div>
              <span className="text-sm font-medium text-gray-700">Administrador</span>
            </div>
            <button
              onClick={logout}
              className="md:hidden p-2 text-gray-500 hover:text-red-600 transition-colors"
            >
              <LogOut className="h-5 w-5" />
            </button>
          </div>
        </header>

        {/* Page Content */}
        <main className="flex-1 p-8 overflow-auto">
          {children}
        </main>
      </div>
    </div>
  );
}
