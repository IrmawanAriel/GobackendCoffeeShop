package models

import "time"

type Product struct {
	Id           int        `db:"id" json:"id"`
	Description  string     `db:"description" json:"description"`
	Category     string     `db:"category" json:"category"`
	Stock        int        `db:"stock" json:"stock"`
	Price        int        `db:"price" json:"price"`
	Rating       *float32   `db:"rating" json:"rating"`
	Product_name string     `db:"product_name" json:"product_name"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at"`
	Image        *string    `db:"image" json:"image"`
	Uuid         string     `db:"uuid" json:"uuid"`
}

type Products []Product

//favorite product

type Favorite struct {
	User_id    int `db:"user_id" json:"user_id"`
	Product_id int `db"product_id" json:"product_id"`
}

type Favorites []Favorite
