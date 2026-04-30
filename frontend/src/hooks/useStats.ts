import { useQuery } from '@tanstack/react-query';
import api from '@/services/api';
import { StatsResponse } from '@/types/stats';

export const useStats = () => {
  return useQuery({
    queryKey: ['stats'],
    queryFn: async () => {
      const response = await api.get<any>('/stats');
      
      // Handle case where response is wrapped in a 'data' field
      if (response.data && response.data.data && !response.data.total_orders) {
        return response.data.data as StatsResponse;
      }
      
      return response.data as StatsResponse;
    },
  });
};
