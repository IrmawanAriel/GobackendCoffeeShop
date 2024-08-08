package repositories

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"fmt"
	"strings"

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
	q := `SELECT id, fullname, email, password, role, address, image, phone FROM public.users`
	data := models.Users{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoUser) UpdateUser(id string, data *models.User) (string, error) {
	q1 := `SELECT id FROM public.users where id = :id`
	check := models.User{}

	if err := r.Select(&check, q1, id); err != nil {
		return "No such a User", err
	}

	setClauses := []string{}
	params := map[string]interface{}{"id": id}

	if data.Fullname != "" {
		setClauses = append(setClauses, "fullname = :fullname")
		params["fullname"] = data.Fullname
	}
	if data.Email != "" {
		setClauses = append(setClauses, "email = :email")
		params["email"] = data.Email
	}
	if data.Password != "" {
		setClauses = append(setClauses, "password = :password")
		params["password"] = data.Password
	}
	if data.Role != "" {
		setClauses = append(setClauses, "role = :role")
		params["role"] = data.Role
	}
	if data.Address != nil {
		setClauses = append(setClauses, "address = :address")
		params["address"] = *data.Address
	}
	if data.Image != nil {
		setClauses = append(setClauses, "image = :image")
		params["image"] = *data.Image
	}
	if data.Phone != nil {
		setClauses = append(setClauses, "phone = :phone")
		params["phone"] = *data.Phone
	}

	if len(setClauses) == 0 {
		return "", fmt.Errorf("no fields to update")
	}

	q := fmt.Sprintf(`UPDATE public.users SET %s WHERE id = :id`, strings.Join(setClauses, ", "))

	_, err := r.DB.NamedExec(q, params)
	if err != nil {
		return "Update Failed", err
	}

	return "User updated successfully", nil
}

func (r *RepoUser) InsertUser(data *models.User) (string, error) {
	q := `INSERT INTO public.users (fullname, email, password)
          VALUES (:fullname, :email, :password)`

	_, err := r.DB.NamedExec(q, data)
	if err != nil {
		return "Create Failed", err
	}

	return "User inserted successfully", nil
}

func (h *RepoUser) DeleteUserById(id int) (string, error) {
	q := `DELETE FROM public.users WHERE id = :id`
	q2 := `DELETE FROM public.favorite_product where user_id = :id`

	params := map[string]interface{}{
		"id": id,
	}

	h.NamedExec(q2, params)

	_, err := h.NamedExec(q, params)
	if err != nil {
		return "Delete Failed", err
	}

	return "User deleted successfully", nil
}
