package models

import "time"

type Product struct {
	Id           string     `db:"id" json:"id"`
	Description  string     `db:"description" json:"description,omitempty" form:"description" valid:"stringlength(5|100)~description minimal 5 dan maksimal 100"`
	Category     string     `db:"category" json:"category" form:"category" valid:"stringlength(4|100)~Category minimal 4 dan maksimal 100"`
	Stock        int        `db:"stock" json:"stock" form:"stock"`
	Price        int        `db:"price" json:"price" form:"price" valid:"range(10000|100000)~Price minimal 10.000 dan maksimal 100.000"`
	Rating       *float32   `db:"rating" json:"rating,omitempty" valid:"range(|5)~Mkasimal rating "`
	Product_name string     `db:"product_name" json:"product_name" form:"product_name" valid:"stringlength(5|100)~Nama produk minimal 5 dan maksimal 100"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Image        *string    `db:"image" json:"image" valid:"stringlength(10|100)~Image is invalid"`
	// Uuid         string     `db:"uuid" json:"uuid"`
}

type Pagination struct {
	Products []Product `json:"products"`
	Page     int
	Limit    int
	Total    int `json:"total"`
}

type Products []Product
