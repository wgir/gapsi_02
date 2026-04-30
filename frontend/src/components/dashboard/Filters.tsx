'use client';

import { useFilterStore } from '@/store/useFilterStore';
import { Search, RotateCcw } from 'lucide-react';

export default function Filters() {
  const { filters, setFilter, resetFilters } = useFilterStore();

  const handleFilterChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFilter(name as any, value);
  };

  return (
    <div className="bg-white p-6 rounded-xl shadow-sm border border-gray-100 mb-8">
      <div className="flex items-center justify-between mb-4">
        <h3 className="text-sm font-semibold text-gray-900 flex items-center gap-2">
          <Search className="h-4 w-4" />
          Filtros Globales
        </h3>
        <button
          onClick={resetFilters}
          className="text-xs text-gray-500 hover:text-blue-600 flex items-center gap-1 transition-colors"
        >
          <RotateCcw className="h-3 w-3" />
          Restablecer
        </button>
      </div>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
        <div>
          <label htmlFor="canal" className="block text-xs font-medium text-gray-500 mb-1">
            Canal
          </label>
          <select
            id="canal"
            name="canal"
            value={filters.canal}
            onChange={handleFilterChange}
            className="block w-full rounded-lg border border-gray-200 py-2.5 px-3 text-sm focus:border-blue-500 focus:ring-blue-500 bg-gray-50"
          >
            <option value="">Todos los canales</option>
            <option value="APP">APP</option>
            <option value="WEB">WEB</option>
            <option value="MOBILE">MOBILE</option>
            <option value="APV AND">APV AND</option>
          </select>
        </div>

        <div>
          <label htmlFor="company" className="block text-xs font-medium text-gray-500 mb-1">
            Compañía
          </label>
          <select
            id="company"
            name="company"
            value={filters.company}
            onChange={handleFilterChange}
            className="block w-full rounded-lg border border-gray-200 py-2.5 px-3 text-sm focus:border-blue-500 focus:ring-blue-500 bg-gray-50"
          >
            <option value="">Todas las compañías</option>
            <option value="BODEGA-01">BODEGA-01</option>
            <option value="LIVERPOOL">LIVERPOOL</option>
          </select>
        </div>

        <div>
          <label htmlFor="fulfillment_type" className="block text-xs font-medium text-gray-500 mb-1">
            Tipo de Entrega
          </label>
          <select
            id="fulfillment_type"
            name="fulfillment_type"
            value={filters.fulfillment_type}
            onChange={handleFilterChange}
            className="block w-full rounded-lg border border-gray-200 py-2.5 px-3 text-sm focus:border-blue-500 focus:ring-blue-500 bg-gray-50"
          >
            <option value="">Todos los tipos</option>
            <option value="FFW">FFW</option>
            <option value="Liverpool_CNC_PICK_PACK">PICK PACK</option>
            <option value="Fulfillment_Type_Liverpool">LIVERPOOL FF</option>
          </select>
        </div>

        <div>
          <label htmlFor="product_type" className="block text-xs font-medium text-gray-500 mb-1">
            Tipo de Producto
          </label>
          <select
            id="product_type"
            name="product_type"
            value={filters.product_type}
            onChange={handleFilterChange}
            className="block w-full rounded-lg border border-gray-200 py-2.5 px-3 text-sm focus:border-blue-500 focus:ring-blue-500 bg-gray-50"
          >
            <option value="">Todos los productos</option>
            <option value="FOOD">FOOD</option>
            <option value="BT">BIG TICKET</option>
            <option value="SL">SOFT LINE</option>
          </select>
        </div>
      </div>
    </div>
  );
}
