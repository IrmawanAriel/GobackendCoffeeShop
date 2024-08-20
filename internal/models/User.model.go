package models

import (
	"time"
)

type User struct {
	Id         string     `db:"id" json:"id"`
	Fullname   string     `db:"fullname" json:"fullname" valid:"required,stringlength(3|100)~Nama lengkap minimal 3 dan maksimal 100 karakter"`
	Email      string     `db:"email" json:"email" valid:"required,email~Format email tidak valid"`
	Password   string     `db:"password" json:"password" valid:"required,stringlength(8|100)~Password minimal 8 karakter"`
	Address    *string    `db:"address" json:"address,omitempty" valid:"stringlength(5|255)~Alamat minimal 5 dan maksimal 255 karakter"`
	Created_at *time.Time `db:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Is_deleted *bool      `db:"is_deleted" json:"is_deleted,omitempty"`
	Image      *string    `db:"image" json:"image,omitempty"`
	Role       string     `db:"role" json:"role" valid:"required,stringlength(3|50)~Role minimal 3 dan maksimal 50 karakter"`
	Phone      *string    `db:"phone" json:"phone,omitempty" valid:"matches(^\\+?[1-9]\\d{1,14}$)~Nomor telepon tidak valid"`
}

type UserCreate struct {
	Id         string     `db:"id" json:"id"`
	Fullname   string     `db:"fullname" json:"fullname" valid:"required,stringlength(3|100)~Nama lengkap minimal 3 dan maksimal 100 karakter"`
	Email      string     `db:"email" json:"email" valid:"required,email~Format email tidak valid"`
	Password   string     `db:"password" json:"password" valid:"required,stringlength(8|100)~Password minimal 8 karakter"`
	Address    *string    `db:"address" json:"address,omitempty" valid:"stringlength(5|255)~Alamat minimal 5 dan maksimal 255 karakter"`
	Created_at *time.Time `db:"created_at" json:"created_at,omitempty"`
	Updated_at *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	Is_deleted *bool      `db:"is_deleted" json:"is_deleted,omitempty"`
	Image      *string    `db:"image" json:"image,omitempty"`
	Phone      *string    `db:"phone" json:"phone,omitempty" valid:"matches(^\\+?[1-9]\\d{1,14}$)~Nomor telepon tidak valid"`
}

type Login struct {
	Email    string `db:"email" json:"email" valid:"required,email~Format email tidak valid"`
	Password string `db:"password" json:"password" valid:"required,stringlength(8|100)~Password minimal 8 karakter"`
}

type Users []User
