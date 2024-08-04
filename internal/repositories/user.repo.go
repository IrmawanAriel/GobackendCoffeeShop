package repositories

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) GetUserById(id string) (models.User, error) {
	var data models.User
	q := `SELECT id, fullname, email, password, role, address, image, phone FROM users WHERE id = $1`
	err := r.Get(&data, q, id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *RepoUser) GetAllUser() (*models.Users, error) {
	q := `SELECT * FROM public.users`
	data := models.Users{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}

// func (r *RepoUser) UpdateUser (string, error) {

// }
