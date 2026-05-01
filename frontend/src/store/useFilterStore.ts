import { create } from 'zustand';
import { FilterParams } from '@/types/stats';

interface FilterState {
  filters: FilterParams;
  setFilter: (name: keyof FilterParams, value: any) => void;
  resetFilters: () => void;
  setPage: (page: number) => void;
}

const initialFilters: FilterParams = {
  canal: '',
  company: '',
  fulfillment_type: '',
  product_type: '',
  page: 1,
  page_size: 10,
};

export const useFilterStore = create<FilterState>((set) => ({
  filters: initialFilters,
  setFilter: (name, value) => 
    set((state) => ({ 
      filters: { ...state.filters, [name]: value, page: 1 } // Reset page on filter change
    })),
  setPage: (page) =>
    set((state) => ({
      filters: { ...state.filters, page }
    })),
  resetFilters: () => set({ filters: initialFilters }),
}));
