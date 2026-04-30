'use client';

import { useStats } from '@/hooks/useStats';
import { useFilterStore } from '@/store/useFilterStore';
import { ShoppingBag, AlertTriangle, TrendingUp, Package } from 'lucide-react';

export default function StatsCards() {
  const { filters } = useFilterStore();
  const { data, isLoading } = useStats(filters);

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {[...Array(4)].map((_, i) => (
          <div key={i} className="bg-white p-6 rounded-xl shadow-sm border border-gray-100 animate-pulse">
            <div className="h-4 w-24 bg-gray-200 rounded mb-4"></div>
            <div className="h-8 w-16 bg-gray-200 rounded"></div>
          </div>
        ))}
      </div>
    );
  }

  const totalOrders = data?.total_orders || 0;
  const errorPercentage = data?.percentage_with_errors || 0;
  const productTypes = Object.keys(data?.breakdown_by_product_type || {}).length;

  const stats = [
    {
      title: 'Total Órdenes',
      value: totalOrders.toLocaleString(),
      icon: ShoppingBag,
      color: 'blue',
      trend: '+12.5%',
    },
    {
      title: 'Órdenes con Error',
      value: `${errorPercentage.toFixed(2)}%`,
      icon: AlertTriangle,
      color: 'red',
      trend: '-2.1%',
    },
    {
      title: 'Promedio Diario',
      value: Math.round(totalOrders / 30).toLocaleString(),
      icon: TrendingUp,
      color: 'green',
      trend: '+4.3%',
    },
    {
      title: 'Tipos de Producto',
      value: productTypes.toString(),
      icon: Package,
      color: 'purple',
      trend: 'Estable',
    },
  ];

  const getColorClasses = (color: string) => {
    const colors: Record<string, string> = {
      blue: 'bg-blue-50 text-blue-600',
      red: 'bg-red-50 text-red-600',
      green: 'bg-green-50 text-green-600',
      purple: 'bg-purple-50 text-purple-600',
    };
    return colors[color] || colors.blue;
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      {stats.map((stat, index) => (
        <div key={index} className="bg-white p-6 rounded-xl shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
          <div className="flex items-center justify-between mb-4">
            <div className={`p-2 rounded-lg ${getColorClasses(stat.color)}`}>
              <stat.icon className="h-5 w-5" />
            </div>
            <span className={`text-xs font-medium ${stat.trend.startsWith('+') ? 'text-green-600' : stat.trend.startsWith('-') ? 'text-red-600' : 'text-gray-500'}`}>
              {stat.trend}
            </span>
          </div>
          <div>
            <p className="text-sm font-medium text-gray-500">{stat.title}</p>
            <p className="text-2xl font-bold text-gray-900 mt-1">{stat.value}</p>
          </div>
        </div>
      ))}
    </div>
  );
}
