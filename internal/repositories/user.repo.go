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

	q := `select id, fullname, email, password, role, address, image, phone from users where id = :id`

	param := map[string]interface{}{
		"id": id,
	}

	err := r.Get(&data, q, param)
	if err != nil {
		return models.User{}, err
	}

	return data, nil

}
