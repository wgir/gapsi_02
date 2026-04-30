'use client';

import { useOrders } from '@/hooks/useOrders';
import { Order } from '@/types/order';
import { useFilterStore } from '@/store/useFilterStore';
import { ChevronLeft, ChevronRight, Eye, AlertCircle } from 'lucide-react';
import { format } from 'date-fns';
import { es } from 'date-fns/locale';

export default function OrdersTable() {
  const { filters, setPage } = useFilterStore();
  const { data: ordersData, isLoading, isError } = useOrders(filters);

  if (isLoading) {
    return (
      <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
        <div className="p-6 border-b border-gray-100">
          <div className="h-5 w-32 bg-gray-200 rounded animate-pulse"></div>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full">
            <thead className="bg-gray-50">
              <tr>
                {[...Array(8)].map((_, i) => (
                  <th key={i} className="px-6 py-4">
                    <div className="h-4 w-20 bg-gray-200 rounded animate-pulse"></div>
                  </th>
                ))}
              </tr>
            </thead>
            <tbody>
              {[...Array(5)].map((_, i) => (
                <tr key={i} className="border-b border-gray-100">
                  {[...Array(8)].map((_, j) => (
                    <td key={j} className="px-6 py-4">
                      <div className="h-4 w-full bg-gray-100 rounded animate-pulse"></div>
                    </td>
                  ))}
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden flex flex-col items-center justify-center text-red-500 py-20">
        <AlertCircle className="h-12 w-12 mb-3 opacity-50" />
        <span className="text-base font-medium">Error al cargar las órdenes</span>
      </div>
    );
  }

  const orders = ordersData?.data || [];
  const totalPages = ordersData?.total_pages || 0;
  const currentPage = ordersData?.current_page || 1;

  const getStatusColor = (status: string) => {
    if (!status) return 'bg-gray-100 text-gray-700';
    switch (status.toUpperCase()) {
      case 'ENTREGADO':
        return 'bg-green-100 text-green-700';
      case 'PENDIENTE_ENVIO':
      case 'PENDIENTE':
        return 'bg-blue-100 text-blue-700';
      case 'PENDIENTE_PAGO':
        return 'bg-yellow-100 text-yellow-700';
      case 'CANCELADO':
      case 'ERROR':
        return 'bg-red-100 text-red-700';
      default:
        return 'bg-gray-100 text-gray-700';
    }
  };

  const parseTimestamp = (timestampStr: string) => {
    if (!timestampStr) return null;
    const cleanDate = timestampStr.replace('__Timestamp__', '');
    try {
      const date = new Date(cleanDate);
      return isNaN(date.getTime()) ? null : date;
    } catch {
      return null;
    }
  };

  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
      <div className="p-6 border-b border-gray-100 flex items-center justify-between">
        <h3 className="text-sm font-semibold text-gray-900">Listado de Órdenes</h3>
        <span className="text-xs text-gray-500">
          Mostrando {orders.length} de {ordersData?.total_rows || 0} resultados
        </span>
      </div>
      
      <div className="overflow-x-auto">
        <table className="w-full text-sm text-left">
          <thead className="bg-gray-50 text-gray-500 uppercase text-[10px] font-bold tracking-wider">
            <tr>
              <th className="px-6 py-4">No. Pedido</th>
              <th className="px-6 py-4">Canal</th>
              <th className="px-6 py-4">Compañía</th>
              <th className="px-6 py-4">Producto</th>
              <th className="px-6 py-4">Fulfillment</th>
              <th className="px-6 py-4">SKU</th>
              <th className="px-6 py-4">Fecha</th>
              <th className="px-6 py-4">Estado</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-100">
            {orders.length === 0 ? (
              <tr>
                <td colSpan={8} className="px-6 py-8 text-center text-gray-500">
                  No se encontraron órdenes.
                </td>
              </tr>
            ) : (
              orders.map((order: Order) => {
                const purchaseDate = parseTimestamp(order.fechaCompra);
                const status = order.orderStatus?.description || (order.error ? 'Error' : 'Pendiente');
                const statusCode = order.orderStatus?.code || (order.error ? 'ERROR' : 'PENDIENTE');

                return (
                  <tr key={order.id} className="hover:bg-gray-50 transition-colors">
                    <td className="px-6 py-4 font-medium text-gray-900">
                      {order.noPedido || order.code || 'N/A'}
                    </td>
                    <td className="px-6 py-4">
                      <span className="px-2 py-1 rounded-md bg-gray-100 text-gray-600 text-[10px] font-medium uppercase">
                        {order.canal}
                      </span>
                    </td>
                    <td className="px-6 py-4 text-gray-600">
                      {order.company}
                    </td>
                    <td className="px-6 py-4">
                      <span className="px-2 py-1 rounded-md bg-blue-50 text-blue-700 text-[10px] font-medium uppercase">
                        {order.productType}
                      </span>
                    </td>
                    <td className="px-6 py-4 text-gray-500 text-xs">
                      {order.fulfillmentType}
                    </td>
                    <td className="px-6 py-4">
                      <span className="text-gray-900 font-medium">{order.sku}</span>
                    </td>
                    <td className="px-6 py-4 text-gray-600 whitespace-nowrap">
                      {purchaseDate 
                        ? format(purchaseDate, 'dd/MM/yy', { locale: es })
                        : 'N/A'}
                    </td>
                    <td className="px-6 py-4">
                      <span className={`px-2.5 py-1 rounded-full text-[10px] font-bold ${getStatusColor(statusCode)}`}>
                        {status}
                      </span>
                    </td>
                  </tr>
                );
              })
            )}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <div className="p-4 border-t border-gray-100 bg-gray-50 flex items-center justify-between">
        <button
          onClick={() => setPage(currentPage - 1)}
          disabled={currentPage === 1}
          className="flex items-center gap-1 px-3 py-1.5 text-sm font-medium text-gray-600 bg-white border border-gray-200 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
        >
          <ChevronLeft className="h-4 w-4" />
          Anterior
        </button>
        
        <div className="flex items-center gap-2">
          {totalPages > 0 && Array.from({ length: Math.min(5, totalPages) }, (_, i) => {
            let pageNum = currentPage;
            if (currentPage <= 3) pageNum = i + 1;
            else if (currentPage >= totalPages - 2) pageNum = totalPages - 4 + i;
            else pageNum = currentPage - 2 + i;
            
            if (pageNum <= 0 || pageNum > totalPages) return null;

            return (
              <button
                key={pageNum}
                onClick={() => setPage(pageNum)}
                className={`h-8 w-8 text-sm font-medium rounded-lg transition-all ${
                  currentPage === pageNum
                    ? 'bg-blue-600 text-white shadow-md'
                    : 'bg-white text-gray-600 border border-gray-200 hover:bg-gray-50'
                }`}
              >
                {pageNum}
              </button>
            );
          })}
        </div>

        <button
          onClick={() => setPage(currentPage + 1)}
          disabled={currentPage === totalPages || totalPages === 0}
          className="flex items-center gap-1 px-3 py-1.5 text-sm font-medium text-gray-600 bg-white border border-gray-200 rounded-lg hover:bg-gray-50 disabled:opacity-50 disabled:cursor-not-allowed transition-all"
        >
          Siguiente
          <ChevronRight className="h-4 w-4" />
        </button>
      </div>
    </div>
  );
}
