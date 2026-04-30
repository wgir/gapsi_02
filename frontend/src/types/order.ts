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
  // Keep some old fields for compatibility
  code?: string;
  orderStatus?: OrderStatus;
  totalPayment?: string;
  orderItems?: OrderItem[];
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
  total_pages?: number; // Optional as we might calculate it
  total_rows?: number;  // Optional as we might calculate it
}

// ... other types stay the same
export interface PaymentStatus {
  code: string;
  description: string;
  id: number;
  can_cancel: boolean;
}

export interface OrderAddress {
  id: number;
  apartment: string | null;
  streetAddress: string;
  addressLine2: string | null;
  postalCode: string | null;
  addressName: string;
  latitude: number | null;
  longitude: number | null;
  phone: string;
  floor: number | null;
  district: string | null;
  city: string;
  state: string;
  country: string;
  neighborhood: string | null;
  reference: string | null;
}

export interface CustomerDetails {
  id: number;
  userId: number;
  documentType: DocumentType;
  documentNumber: string;
  name: string;
  email: string;
}

export interface DocumentType {
  id: number;
  code: string;
  description: string;
}

export interface PaymentMethod {
  id: number;
  code: string;
  description: string;
  paymentMethodType: PaymentMethodType;
  isOnline: boolean;
  canCancel: boolean;
}

export interface PaymentMethodType {
  id: number;
  code: string;
  description: string;
}

export interface DeliveryMethod {
  id: number;
  code: string;
  description: string;
  canCancel: boolean;
}

export interface OrderStatusHistory {
  id: number;
  status: OrderStatus;
  isCurrent: boolean;
  dateCreated: string;
  userEmail: string;
}

export interface OrderItem {
  id: number;
  orderCode: string;
  orderId: number;
  sku: string;
  name: string;
  quantity: number;
  unitPrice: string;
  unitPriceWithTaxes: string;
  lineTotal: string;
  lineTotalWithTaxes: string;
  isCombo: boolean;
  comboItems?: ComboItem[];
}

export interface ComboItem {
  id: number;
  name: string;
  sku: string;
  quantity: number;
  isComboItem: boolean;
  unitPrice: string;
  unitPriceWithTaxes: string;
  lineTotal: string;
  lineTotalWithTaxes: string;
}
