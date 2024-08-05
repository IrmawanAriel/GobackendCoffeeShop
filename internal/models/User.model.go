package models

import (
	"time"
)

type User struct {
	Id         int        `db:"id" json:"id"`
	Fullname   string     `db:"fullname" json:"fullname"`
	Email      string     `db:"email" json:"email"`
	Password   string     `db:"password" json:"password"`
	Address    *string    `db:"address" json:"address"`
	Created_at *time.Time `db:"created_at" json:"created_at"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at"`
	Is_deleted *bool      `db:"is_deleted" json:"is_deleted"`
	Image      *string    `db:"image" json:"image"`
	Role       string     `db:"role" json:"role"`
	Phone      *string    `db:"phone" json:"phone"`
}

type Users []User
