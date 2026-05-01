import { useQuery } from '@tanstack/react-query';
import { api } from '@/services/api';

export interface FilterOptions {
  channels: string[];
  companies: string[];
  fulfillmentTypes: string[];
  productTypes: string[];
}

export function useFilters() {
  return useQuery<FilterOptions>({
    queryKey: ['orderFilters'],
    queryFn: async () => {
      const { data } = await api.get<FilterOptions>('/filters');
      return data;
    },
    staleTime: 1000 * 60 * 30, // 30 minutes
  });
}
