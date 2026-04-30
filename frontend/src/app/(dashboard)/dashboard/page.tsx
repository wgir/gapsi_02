'use client';

import Filters from '@/components/dashboard/Filters';
import StatsCards from '@/components/dashboard/StatsCards';
import OrdersCharts from '@/components/dashboard/OrdersCharts';
import OrdersTable from '@/components/dashboard/OrdersTable';

export default function DashboardPage() {
  return (
    <div className="max-w-7xl mx-auto space-y-8">
      <div>
        <h2 className="text-2xl font-bold text-gray-900">Dashboard de Órdenes</h2>
        <p className="text-sm text-gray-500 mt-1">
          Visualiza y gestiona las órdenes de venta en tiempo real.
        </p>
      </div>

      <Filters />
      
      <StatsCards />

      <OrdersCharts />

      <OrdersTable />
    </div>
  );
}
