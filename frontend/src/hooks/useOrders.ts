import { useQuery } from '@tanstack/react-query';
import api from '@/services/api';
import { OrdersResponse } from '@/types/order';
import { FilterParams } from '@/types/stats';
import { normaliseOrdersPage } from '@/lib/adapters/order';

export const useOrders = (filters: FilterParams) => {
  return useQuery({
    queryKey: ['orders', filters],
    queryFn: async () => {
      const { data } = await api.post<OrdersResponse>('/orders', filters);
      return normaliseOrdersPage(data, filters);
    },
    staleTime: 1000 * 30, // 30 seconds — orders change frequently
    placeholderData: (prev) => prev, // keep previous data while loading next page
  });
};
