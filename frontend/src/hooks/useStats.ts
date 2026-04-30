import { useQuery } from '@tanstack/react-query';
import api from '@/services/api';
import { StatsResponse, FilterParams } from '@/types/stats';

export const useStats = (filters: FilterParams) => {
  return useQuery({
    queryKey: ['stats', filters],
    queryFn: async () => {
      const params = new URLSearchParams();
      if (filters.canal) params.append('canal', filters.canal);
      if (filters.company) params.append('company', filters.company);
      if (filters.fulfillment_type) params.append('fulfillment_type', filters.fulfillment_type);
      if (filters.product_type) params.append('product_type', filters.product_type);

      const response = await api.get<any>(`/stats?${params.toString()}`);
      
      // Handle case where response is wrapped in a 'data' field
      if (response.data && response.data.data && !response.data.total_orders) {
        return response.data.data as StatsResponse;
      }
      
      return response.data as StatsResponse;
    },
  });
};
