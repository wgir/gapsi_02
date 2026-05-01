// Server Component — no 'use client' needed here
import { User } from 'lucide-react';
import SidebarClient from '@/components/layout/SidebarClient';
import LogoutButtonClient from '@/components/layout/LogoutButtonClient';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex min-h-screen bg-gray-50">
      <SidebarClient />

      {/* Main Content */}
      <div className="flex-1 flex flex-col">
        {/* Top Header */}
        <header className="h-16 bg-white border-b border-gray-200 flex items-center justify-between px-8">
          <div className="md:hidden">
            <h1 className="text-xl font-bold text-blue-600">Gapsi</h1>
          </div>
          <div className="flex-1 hidden md:block" />
          <div className="flex items-center gap-4">
            <div className="flex items-center gap-2 px-3 py-1.5 bg-gray-50 rounded-full border border-gray-100">
              <div className="h-8 w-8 rounded-full bg-blue-100 flex items-center justify-center">
                <User className="h-5 w-5 text-blue-600" />
              </div>
              <span className="text-sm font-medium text-gray-700">Administrador</span>
            </div>
            {/* Mobile logout */}
            <LogoutButtonClient />
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
