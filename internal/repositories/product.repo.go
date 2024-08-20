// INSERT INTO "product"("description","category","stock","price","rating","product_name") VALUES('test','coffee',3,25000,3.5,'Matcha Latte');
package repositories

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductRepositoryInterface interface {
	CreateProduct(data *models.Product) (string, error)
	GetAllProduct(search string, sort string, category string, pagination *models.Pagination) (*models.Products, error)
	GetProductById(id string) (*models.Product, error)
	UpdateProduct(id string, data *models.Product) (string, error)
	DeleteProductById(id string) (string, error)
	GetFavoritesProduct(userID string) (*models.Products, string, error)
	AddFavoriteProduct(userId string, productId string) (string, error)
	DeleteFavoriteProduct(userId string, productId string) (string, error)
}

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {
	q := `INSERT INTO public.product ("description","category","stock","price","rating","product_name","image") 
	VALUES(
	:description,
	:category,
	:stock,
	:price,
	:rating,
	:product_name,
	:image )`

	_, err := r.NamedExec(q, data)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return "error", fmt.Errorf("product name already exists: %w", err)
		}
		return "error in", fmt.Errorf("query execution error: %w", err)
	}

	return "1 data product created", nil
}

func (r *RepoProduct) GetAllProduct(search string, sort string, category string, pagination *models.Pagination) (*models.Products, error) {
	baseQuery := `SELECT * FROM public.product`
	var conditions []string
	params := make(map[string]interface{})

	if search != "" {
		conditions = append(conditions, "product_name ILIKE :search")
		params["search"] = "%" + search + "%"
	}

	if category != "" {
		conditions = append(conditions, "category = :category")
		params["category"] = category
	}

	if len(conditions) > 0 {
		baseQuery += " WHERE " + strings.Join(conditions, " AND ")
	}

	if sort != "" {
		baseQuery += " ORDER BY " + sort
	}

	if pagination.Limit > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		baseQuery += " LIMIT :limit OFFSET :offset "
		params["limit"] = pagination.Limit
		params["offset"] = offset
	}

	data := models.Products{}
	query, args, err := sqlx.Named(baseQuery, params)
	if err != nil {
		return nil, err
	}

	query = r.Rebind(query)

	if err := r.Select(&data, query, args...); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoProduct) GetProductById(id string) (*models.Product, error) {
	q := "SELECT * FROM public.product WHERE id = :id"
	var data models.Product

	if err := r.Select(&data, q); err != nil {
		return nil, fmt.Errorf("query execution error: %s", id)
	}

	return &data, nil
}

func (r *RepoProduct) UpdateProduct(id string, data *models.Product) (string, error) {
	q := `UPDATE public.product
		SET description = :description,
			category = :category,
			stock = :stock,
			price = :price,
			rating = :rating,
			product_name = :product_name
		WHERE id = :id
		`
	params := map[string]interface{}{
		"id":           id,
		"description":  data.Description,
		"category":     data.Category,
		"stock":        data.Stock,
		"price":        data.Price,
		"rating":       data.Rating,
		"product_name": data.Product_name,
		"image":        data.Image,
	}

	_, err := r.NamedExec(q, params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return "error", fmt.Errorf("product name already exists: %w", err)
		}
		return "error in", fmt.Errorf("query execution error: %w", err)
	}

	return "Product updated successfully", nil
}

func (r *RepoProduct) DeleteProductById(id string) (string, error) {
	q := `DELETE FROM public.product WHERE id = :id`
	q2 := `DELETE FROM public.favorite_product WHERE product_id = :id`

	params := map[string]interface{}{
		"id": id,
	}

	r.NamedExec(q2, params)

	_, err := r.NamedExec(q, params)
	if err != nil {
		return "Delete Failed", err
	}

	return "Product deleted successfully", nil

}

func (r *RepoProduct) GetFavoritesProduct(userID string) (*models.Products, string, error) {
	q := `SELECT p.*
          FROM product p
          JOIN favorite_product fp ON p.id = fp.product_id
          WHERE fp.user_id = $1;
          `
	var products models.Products

	err := r.Select(&products, q, userID)
	if err != nil {
		return nil, " Massage: Product not found ", err
	}

	return &products, "", nil
}

func (r RepoProduct) AddFavoriteProduct(userId string, productId string) (string, error) {
	q := `INSERT INTO public.favorite_product ("user_id","product_id") 
	VALUES(
	:user_id,
	:product_id )`

	params := map[string]interface{}{
		"user_id":    userId,
		"product_id": productId,
	}

	_, err := r.NamedExec(q, params)
	if err != nil {
		return "", fmt.Errorf("no product with id: %s\nor no user with id : %s", productId, userId)
	}

	return "Product added to favorite successfully", nil

}

func (r RepoProduct) DeleteFavoriteProduct(userId string, productId string) (string, error) {
	q := `DELETE FROM public.favorite_product WHERE user_id = :user_id and product_id = :product_id`

	params := map[string]interface{}{
		"user_id":    userId,
		"product_id": productId,
	}

	_, err := r.NamedExec(q, params)
	if err != nil {
		return "Delete Failed, no such product", err
	}

	return "Product deleted from favorite successfully", nil

}
