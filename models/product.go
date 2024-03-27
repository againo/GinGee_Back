package models

// models/product.go

package models

import "time"

// Product represents a product model.
type Product struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Barcode      string    `json:"barcode"`
	StockQuantity int       `json:"stock_quantity"`
	Unit         string    `json:"unit"`
	ShelfLife    int       `json:"shelf_life"`
	ProductionDate time.Time `json:"production_date"`
	Image        string    `json:"image"`
	Category     string    `json:"category"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Status       string    `json:"status"`
	Featured     bool      `json:"featured"`
	Visible      bool      `json:"visible"`
	DateAdded    time.Time `json:"date_added"`
}
