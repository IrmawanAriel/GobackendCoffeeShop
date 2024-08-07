package models

import (
	"time"
)

type User struct {
	Id         int        `db:"id" json:"id"` //
	Fullname   string     `db:"fullname" json:"fullname"`
	Email      string     `db:"email" json:"email"`
	Password   string     `db:"password" json:"password"`
	Address    *string    `db:"address" json:"address,omitempty"`
	Created_at *time.Time `db:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Is_deleted *bool      `db:"is_deleted" json:"is_deleted,omitempty"`
	Image      *string    `db:"image" json:"image,omitempty"`
	Role       string     `db:"role" json:"role"`
	Phone      *string    `db:"phone" json:"phone,omitempty"`
}

type Users []User
