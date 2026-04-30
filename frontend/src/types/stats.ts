export interface StatsResponse {
  total_orders: number;
  breakdown_by_canal: Record<string, number>;
  breakdown_by_fulfillment: Record<string, number>;
  breakdown_by_product_type: Record<string, number>;
  percentage_with_errors: number;
}

export interface FilterParams {
  canal?: string;
  company?: string;
  fulfillment_type?: string;
  product_type?: string;
  page?: number;
  page_size?: number;
}
