CREATE TYPE user_role AS ENUM ('ADMIN', 'USER');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role user_role NOT NULL DEFAULT 'USER',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE orders (
    id TEXT PRIMARY KEY,
    canal TEXT NOT NULL,
    cantidad INTEGER NOT NULL,
    company TEXT NOT NULL,
    cp TEXT NOT NULL,
    created_at TEXT NOT NULL,
    days_to_delivery TEXT,
    error TEXT,
    error_message TEXT,
    fecha_compra TEXT,
    fecha_estimada TEXT,
    fulfillment_type TEXT,
    is_flash BOOLEAN NOT NULL DEFAULT FALSE,
    is_marketplace BOOLEAN NOT NULL DEFAULT FALSE,
    no_pedido TEXT,
    plan TEXT,
    product_type TEXT,
    sku TEXT,
    store_selected TEXT,
    tipo_pago TEXT,
    edd1 TEXT,
    edd2 TEXT
);

CREATE INDEX idx_orders_canal ON orders(canal);
CREATE INDEX idx_orders_company ON orders(company);
CREATE INDEX idx_orders_fulfillment_type ON orders(fulfillment_type);
CREATE INDEX idx_orders_product_type ON orders(product_type);
