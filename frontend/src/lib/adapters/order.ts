import { FilterParams } from '@/types/stats';
import { NormalisedOrdersPage, OrdersResponse } from '@/types/order';

/**
 * Transforms the raw API response from /v1/orders into the
 * normalised shape consumed by the UI (OrdersTable, pagination).
 */
export function normaliseOrdersPage(
  raw: OrdersResponse,
  filters: FilterParams,
): NormalisedOrdersPage {
  const totalRows = raw.total ?? (Array.isArray(raw.data) ? raw.data.length : 0);
  const pageSize = raw.page_size ?? filters.page_size ?? 10;

  return {
    data: Array.isArray(raw.data) ? raw.data : [],
    total_rows: totalRows,
    total_pages: Math.ceil(totalRows / pageSize),
    current_page: raw.page ?? filters.page ?? 1,
    page_size: pageSize,
  };
}
