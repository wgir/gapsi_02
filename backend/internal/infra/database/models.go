package database

import (
	"database/sql"
	"time"
)

type UserRole string

const (
	UserRoleADMIN UserRole = "ADMIN"
	UserRoleUSER  UserRole = "USER"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         UserRole  `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Order struct {
	ID              string         `json:"id"`
	Canal           string         `json:"canal"`
	Cantidad        int32          `json:"cantidad"`
	Company         string         `json:"company"`
	Cp              string         `json:"cp"`
	CreatedAt       string         `json:"created_at"`
	DaysToDelivery  sql.NullString `json:"days_to_delivery"`
	Error           sql.NullString `json:"error"`
	ErrorMessage    sql.NullString `json:"error_message"`
	FechaCompra     sql.NullString `json:"fecha_compra"`
	FechaEstimada   sql.NullString `json:"fecha_estimada"`
	FulfillmentType sql.NullString `json:"fulfillment_type"`
	IsFlash         bool           `json:"is_flash"`
	IsMarketplace   bool           `json:"is_marketplace"`
	NoPedido        sql.NullString `json:"no_pedido"`
	Plan            sql.NullString `json:"plan"`
	ProductType     sql.NullString `json:"product_type"`
	Sku             sql.NullString `json:"sku"`
	StoreSelected   sql.NullString `json:"store_selected"`
	TipoPago        sql.NullString `json:"tipo_pago"`
	Edd1            sql.NullString `json:"edd1"`
	Edd2            sql.NullString `json:"edd2"`
}
