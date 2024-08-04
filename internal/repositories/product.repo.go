// INSERT INTO "product"("description","category","stock","price","rating","product_name") VALUES('test','coffee',3,25000,3.5,'Matcha Latte');
package repositories

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

func (r *RepoProduct) CreateProduct(data *models.Product) (string, error) {
	q := `INSERT INTO public.product ("description","category","stock","price","rating","product_name") 
	VALUES(
	:description,
	:category,
	:stock,
	:price,
	:rating,
	:product_name )`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data product created", nil
}

func (r *RepoProduct) GetAllProduct() (*models.Products, error) {
	q := `SELECT * FROM public.product`
	data := models.Products{}

	if err := r.Select(&data, q); err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *RepoProduct) GetProductById(id string, data *models.Product) (string, error) {
	q := `UPDATE public.product
		SET description = :description,
			category = :category,
			stock = :stock,
			price = :price,
			rating = :rating,
			product_name = :product_name
		WHERE id = 1
		`
	params := map[string]interface{}{
		"id":           id,
		"description":  data.Description,
		"category":     data.Category,
		"stock":        data.Stock,
		"price":        data.Price,
		"rating":       data.Rating,
		"product_name": data.Product_name,
	}

	_, err := r.NamedExec(q, params)
	if err != nil {
		return "", err
	}

	return "Product updated successfully", nil
}

func (r *RepoProduct) DeleteProductById(id string) (string, error) {
	q := `DELETE FROM public.product WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}

	_, err := r.NamedExec(q, params)
	if err != nil {
		return "", err
	}

	return "Product updated successfully", nil

}
