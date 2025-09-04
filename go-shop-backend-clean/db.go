package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func getPostgresConnString() string {
	// Check for DATABASE_URL environment variable first
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		return dbURL
	}
	// Build from individual environment variables
	host := getenv("DB_HOST", "localhost")
	port := getenv("DB_PORT", "5432")
	user := getenv("DB_USER", "postgres")
	password := getenv("DB_PASSWORD", "")
	dbname := getenv("DB_NAME", "postgres")
	sslmode := getenv("DB_SSLMODE", "disable")
	
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func connectPostgres() error {
	connStr := getPostgresConnString()
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(2 * time.Hour)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}

func migratePostgres() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			email TEXT NOT NULL,
			name TEXT NOT NULL,
			role TEXT NOT NULL,
			password_hash TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			price NUMERIC(10,2) NOT NULL,
			category TEXT NOT NULL,
			image TEXT NOT NULL,
			stock INT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS orders (
			id UUID PRIMARY KEY,
			order_number TEXT UNIQUE NOT NULL,
			customer_name TEXT NOT NULL,
			customer_email TEXT NOT NULL,
			customer_address TEXT NOT NULL,
			subtotal NUMERIC(10,2) NOT NULL,
			shipping NUMERIC(10,2) NOT NULL,
			tax NUMERIC(10,2) NOT NULL,
			total NUMERIC(10,2) NOT NULL,
			status TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS order_items (
			id SERIAL PRIMARY KEY,
			order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
			product_id INT NOT NULL REFERENCES products(id),
			name TEXT NOT NULL,
			price NUMERIC(10,2) NOT NULL,
			quantity INT NOT NULL
		)`,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}
	return nil
}

func seedPostgresIfEmpty() error {
	// Seed products
	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM products`).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		for _, p := range mockProducts {
			_, err := db.Exec(`INSERT INTO products (name, description, price, category, image, stock) VALUES ($1,$2,$3,$4,$5,$6)`,
				p.Name, p.Description, p.Price, p.Category, p.Image, p.Stock,
			)
			if err != nil {
				return err
			}
		}
		log.Printf("Seeded %d products into Postgres", len(mockProducts))
	}

	// Seed users
	if err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count); err != nil {
		return err
	}
	if count == 0 {
		seedUsers := []struct {
			Username string
			Email    string
			Name     string
			Role     string
			Password string
		}{
			{"admin", "admin@vueshop.com", "Admin User", "admin", "admin123"},
			{"john", "john@example.com", "John Doe", "customer", "password123"},
			{"jane", "jane@example.com", "Jane Smith", "customer", "password456"},
		}
		for _, u := range seedUsers {
			hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			_, err = db.Exec(`INSERT INTO users (username, email, name, role, password_hash) VALUES ($1,$2,$3,$4,$5)`,
				u.Username, u.Email, u.Name, u.Role, string(hash),
			)
			if err != nil {
				return err
			}
		}
		log.Printf("Seeded %d users into Postgres", 3)
	}
	return nil
}
