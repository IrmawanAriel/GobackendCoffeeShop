package seed

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"log"

	"github.com/jmoiron/sqlx"
)

func strPtr(s string) *string {
	return &s
}

func SeedUsers(db *sqlx.DB) error {
	users := []models.User{
		{
			Fullname: "hai",
			Phone:    strPtr("123456789"),
			Address:  strPtr("123 Main St"),
			Email:    "haihaihai@example.com",
			Password: "haihaihai",
			Role:     "admin",
		},
		{
			Fullname: "admin",
			Phone:    strPtr("123456788"),
			Address:  strPtr("123 Main St"),
			Email:    "admin@example.com",
			Password: "haihaihai",
			Role:     "admin",
		},
		{
			Fullname: "ariel",
			Phone:    strPtr("123456787"),
			Address:  strPtr("123 Main St"),
			Email:    "ariel@gmail.com",
			Password: "haihaihai",
			Role:     "admin",
		},
		{
			Fullname: "user",
			Phone:    strPtr("123456787"),
			Address:  strPtr("123 Main St"),
			Email:    "user@gmail.com",
			Password: "haihaihai",
			Role:     "user",
		},
	}

	query := `INSERT INTO public.users(
		fullname,
		phone,
		address,
		email,
		password,
		role
	) VALUES(
		:fullname,
		:phone,
		:address,
		:email,
		:password,
		:role
	)`

	for _, user := range users {
		hashedPassword, err := pkg.HashPassword(user.Password)
		if err != nil {
			return err
		}

		params := map[string]interface{}{
			"fullname": user.Fullname,
			"phone":    user.Phone,
			"address":  user.Address,
			"email":    user.Email,
			"password": hashedPassword,
			"role":     user.Role,
		}

		_, err = db.NamedExec(query, params)
		if err != nil {
			return err
		}
	}

	log.Println("Seeding users completed successfully.")
	return nil
}
