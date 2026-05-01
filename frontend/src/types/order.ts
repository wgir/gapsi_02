// Core order interface — only fields used by the current UI
export interface Order {
  id: string;
  noPedido: string;
  canal: string;
  cantidad: number;
  company: string;
  sku: string;
  fechaCompra: string;
  fechaEstimada: string;
  fulfillmentType: string;
  productType: string;
  createdAt?: string;
  daysToDelivery?: string;
  error?: string;
  errorMessage?: string;
  isFlash?: boolean;
  isMarketplace?: boolean;
  plan?: string;
  // Legacy fields kept for backward-compat with old API responses
  code?: string;
  orderStatus?: OrderStatus;
}

export interface OrderStatus {
  code: string;
  description: string;
  id: number;
  can_cancel: boolean;
}

export interface OrdersResponse {
  data: Order[];
  total: number;
  page: number;
  page_size: number;
  total_pages?: number;
  total_rows?: number;
}

// Normalised shape returned by useOrders (after adapter transform)
export interface NormalisedOrdersPage {
  data: Order[];
  total_rows: number;
  total_pages: number;
  current_page: number;
  page_size: number;
}
