package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/user/gapsi_orders_api/internal/infra/database/sqlc"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	if dbHost == "" { dbHost = "localhost" }
	if dbPort == "" { dbPort = "5432" }
	if dbUser == "" { dbUser = "root" }
	if dbPass == "" { dbPass = "secret" }
	if dbName == "" { dbName = "orders_db" }
	if dbSSL == "" { dbSSL = "disable" }

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPass, dbName, dbSSL)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	// Execute migrations
	schemaBytes, err := os.ReadFile("db/migrations/000001_init_schema.up.sql")
	if err != nil {
		log.Fatalf("failed to read migration file: %v", err)
	}

	fmt.Println("Running schema migration...")
	if _, err := db.Exec(string(schemaBytes)); err != nil {
		log.Fatalf("failed to execute migration: %v", err)
	}
	fmt.Println("Schema migration completed successfully.")

	queries := sqlc.New(db)

	file, err := os.Open("docs/orders_db.csv")
	if err != nil {
		log.Fatalf("failed to open csv: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	// Skip header
	if _, err := reader.Read(); err != nil {
		log.Fatalf("failed to read header: %v", err)
	}

	ctx := context.Background()
	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error reading record: %v", err)
			continue
		}

		// Mapping CSV to DB
		// __id__,canal,cantidad,company,cp,createdAt,daysToDelivery,error,errorMessage,fechaCompra,fechaEstimada,fulfillmentType,isFlash,isMarketplace,noPedido,plan,productType,sku,storeSelected,tipoPago,edd1,edd2

		cantidad, _ := strconv.Atoi(record[2])
		isFlash := record[12] == "TRUE"
		isMarketplace := record[13] == "TRUE"

		err = queries.CreateOrder(ctx, sqlc.CreateOrderParams{
			ID:              record[0],
			Canal:           record[1],
			Cantidad:        int32(cantidad),
			Company:         record[3],
			Cp:              record[4],
			CreatedAt:       record[5],
			DaysToDelivery:  sql.NullString{String: record[6], Valid: record[6] != ""},
			Error:           sql.NullString{String: record[7], Valid: record[7] != ""},
			ErrorMessage:    sql.NullString{String: record[8], Valid: record[8] != ""},
			FechaCompra:     sql.NullString{String: record[9], Valid: record[9] != ""},
			FechaEstimada:   sql.NullString{String: record[10], Valid: record[10] != ""},
			FulfillmentType: sql.NullString{String: record[11], Valid: record[11] != ""},
			IsFlash:         isFlash,
			IsMarketplace:   isMarketplace,
			NoPedido:        sql.NullString{String: record[14], Valid: record[14] != ""},
			Plan:            sql.NullString{String: record[15], Valid: record[15] != ""},
			ProductType:     sql.NullString{String: record[16], Valid: record[16] != ""},
			Sku:             sql.NullString{String: record[17], Valid: record[17] != ""},
			StoreSelected:   sql.NullString{String: record[18], Valid: record[18] != ""},
			TipoPago:        sql.NullString{String: record[19], Valid: record[19] != ""},
			Edd1:            sql.NullString{String: record[20], Valid: record[20] != ""},
			Edd2:            sql.NullString{String: record[21], Valid: record[21] != ""},
		})

		if err != nil {
			log.Printf("failed to insert order %s: %v", record[0], err)
		} else {
			count++
			if count%1000 == 0 {
				fmt.Printf("Inserted %d orders...\n", count)
			}
		}
	}

	fmt.Printf("Seeding completed. Total inserted: %d\n", count)
}
