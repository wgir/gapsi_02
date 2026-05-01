import { create } from 'zustand';
import { FilterParams } from '@/types/stats';

interface FilterState {
  filters: FilterParams;
  // Generic key-constrained setter — no `any` leakage
  setFilter: <K extends keyof FilterParams>(name: K, value: FilterParams[K]) => void;
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
      filters: { ...state.filters, [name]: value, page: 1 }, // reset to page 1 on filter change
    })),
  setPage: (page) =>
    set((state) => ({
      filters: { ...state.filters, page },
    })),
  resetFilters: () => set({ filters: initialFilters }),
}));
