import { useQuery } from '@tanstack/react-query';
import api from '@/services/api';
import { StatsResponse } from '@/types/stats';

export const useStats = () => {
  return useQuery<StatsResponse>({
    queryKey: ['stats'],
    queryFn: async () => {
      const { data } = await api.get<StatsResponse>('/stats');
      return data;
    },
    staleTime: 1000 * 60, // 1 minute — stats are less volatile than order rows
  });
};
