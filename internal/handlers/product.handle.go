package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type HandlerProduct struct {
	repositories.ProductRepositoryInterface
	pkg.CloudinaryInterface
}

func NewProduct(r repositories.ProductRepositoryInterface, cld pkg.CloudinaryInterface) *HandlerProduct {
	return &HandlerProduct{r, cld}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil {
		response.BadRequest("Failed to create product", "Invalid data type")
		return
	}

	_, err := govalidator.ValidateStruct(&product)
	if err != nil {
		response.BadRequest("Failed to create product", err.Error())
		return
	}

	file, header, _ := ctx.Request.FormFile("image")

	if file != nil {
		const maxFileSize = 5 * 1024 * 1024
		if header.Size > maxFileSize {
			response.BadRequest("Failed to create product", "File size exceeds the 5MB limit")
			return
		}

		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			response.BadRequest("Failed to create product", "Only JPG and PNG files are allowed")
			return
		}

		randomNumber := rand.Int()
		fileName := fmt.Sprintf("go-product-%d", randomNumber)
		uploadResult, err := h.UploadFile(ctx, file, fileName)
		if err != nil {
			response.InternalServerError("Failed to upload file", err.Error())
			return
		}

		picture := uploadResult.SecureURL
		product.Image = &picture
	}

	responseData, err := h.CreateProduct(&product)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			response.BadRequest("Failed to create product", "Product name already exists")
			return
		}
		response.InternalServerError("Failed to create product", err.Error())
		return
	}

	response.Created("Product successfully created", responseData)
}

func (h *HandlerProduct) FetchAll(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)

	search := ctx.Query("search")
	sort := ctx.Query("sort")
	category := ctx.Query("category")
	limit := ctx.Query("limit")
	page := ctx.Query("page")

	limits, _ := strconv.Atoi(limit)
	pages, _ := strconv.Atoi(page)

	params := &models.Pagination{
		Limit: limits,
		Page:  pages,
	}

	data, err := h.GetAllProduct(search, sort, category, params)
	if err != nil {
		response.InternalServerError("Failed to retrieve products", err.Error())
		return
	}

	if len(*data) == 0 {
		response.NotFound("No products found", nil)
		return
	}

	response.Success("Successfully retrieved products", data)
}

func (h *HandlerProduct) FetchById(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	idParam := ctx.Param("id")

	data, err := h.GetProductById(idParam)
	if err != nil {
		response.NotFound("Product not found", err.Error())
		return
	}

	response.Success("Successfully retrieved product", data)
}

func (h *HandlerProduct) UpdateById(ctx *gin.Context) {
	product := models.Product{}
	idParam := ctx.Param("id")
	response := pkg.NewResponse(ctx)

	if err := ctx.ShouldBind(&product); err != nil {
		response.BadRequest("Failed to update product", "Invalid data type")
		return
	}

	_, err := govalidator.ValidateStruct(&product)
	if err != nil {
		response.BadRequest("Failed to update product", err.Error())
		return
	}

	file, header, _ := ctx.Request.FormFile("image")

	if file != nil {
		const maxFileSize = 5 * 1024 * 1024
		if header.Size > maxFileSize {
			response.BadRequest("Failed to update product", "File size exceeds the 5MB limit")
			return
		}

		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			response.BadRequest("Failed to update product", "Only JPG and PNG files are allowed")
			return
		}

		randomNumber := rand.Int()
		fileName := fmt.Sprintf("go-product-%d", randomNumber)
		uploadResult, err := h.UploadFile(ctx, file, fileName)
		if err != nil {
			response.InternalServerError("Failed to upload file", err.Error())
			return
		}

		picture := uploadResult.SecureURL
		product.Image = &picture
	}

	data, err := h.UpdateProduct(idParam, &product)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			response.BadRequest("Failed to update product", "Product name already exists")
			return
		}
		response.InternalServerError("Failed to update product", err.Error())
		return
	}

	response.Success("Product successfully updated", data)
}

func (h *HandlerProduct) DeleteProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	data, err := h.DeleteProductById(id)
	if err != nil {
		response.BadRequest("Failed to delete product", "Invalid product ID")
		return
	}

	response.Success("Product successfully deleted", data)
}

func (h *HandlerProduct) GetFavorite(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	user := ctx.Param("userId")

	data, str, err := h.GetFavoritesProduct(user)
	if err != nil {
		response.BadRequest("Failed to retrieve favorite products", str)
		return
	}

	if len(*data) == 0 {
		response.NotFound("No favorite products found", nil)
		return
	}

	response.Success("Successfully retrieved favorite products", data)
}

func (h *HandlerProduct) AddFavorite(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	user := ctx.Param("userId")
	id := ctx.Param("productId")

	data, err := h.AddFavoriteProduct(user, id)
	if err != nil {
		response.BadRequest("Failed to add product to favorites", err.Error())
		return
	}

	response.Success("Product successfully added to favorites", data)
}

func (h *HandlerProduct) DeleteFavorite(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	user := ctx.Param("userId")
	id := ctx.Param("productId")

	data, err := h.DeleteFavoriteProduct(user, id)
	if err != nil {
		response.BadRequest("Failed to remove product from favorites", err.Error())
		return
	}

	response.Success("Product successfully removed from favorites", data)
}
