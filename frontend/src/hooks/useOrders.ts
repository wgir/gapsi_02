import { useQuery } from '@tanstack/react-query';
import api from '@/services/api';
import { OrdersResponse } from '@/types/order';
import { FilterParams } from '@/types/stats';

export const useOrders = (filters: FilterParams) => {
  return useQuery({
    queryKey: ['orders', filters],
    queryFn: async () => {
      const { data } = await api.post<any>('/orders', filters);
      
      // Normalized structure for the UI
      const totalRows = data.total || (Array.isArray(data.data) ? data.data.length : 0);
      const pageSize = data.page_size || filters.page_size || 10;
      const totalPages = Math.ceil(totalRows / pageSize);
      
      return {
        data: Array.isArray(data.data) ? data.data : [],
        total_rows: totalRows,
        total_pages: totalPages,
        current_page: data.page || filters.page || 1,
        page_size: pageSize,
        total: totalRows
      };
    },
  });
};
