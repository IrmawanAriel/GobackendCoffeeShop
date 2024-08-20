package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type HandlerProduct struct {
	repositories.ProductRepositoryInterface
	pkg.Cloudinary
}

func NewProduct(r repositories.ProductRepositoryInterface, cld pkg.Cloudinary) *HandlerProduct {
	return &HandlerProduct{r, cld}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil { // check data type
		response.BadRequest("create data failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&product)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	// get file from request body
	file, header, _ := ctx.Request.FormFile("image")

	if file != nil {
		const maxFileSize = 5 * 1024 * 1024 // 5MB size limit
		if header.Size > maxFileSize {
			response.BadRequest("create data failed, upload file failed, file too large", "file size exceeds the 5MB limit")
			return
		}

		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			response.BadRequest("Create data failed, upload file failed", "Only jpg and png files are allowed")
			return
		}

		randomNumber := rand.Int()
		fileName := fmt.Sprintf("go-product-%d", randomNumber)
		uploadResult, err := h.UploadFile(ctx, file, fileName)
		if err != nil {
			response.BadRequest("create data failed, upload file failed", err.Error())
			return
		}

		picture := uploadResult.SecureURL
		product.Image = &picture
	}

	// create the product
	respone, err := h.CreateProduct(&product)
	if err != nil {
		// Check if the error is related to unique constraint
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product name already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	ctx.JSON(200, respone)
}

func (h *HandlerProduct) FetchAll(ctx *gin.Context) {

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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if len(*data) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No product found"})
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerProduct) FetchById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	data, err := h.GetProductById(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No such product"})
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerProduct) UpdateById(ctx *gin.Context) {
	product := models.Product{}
	idParam := ctx.Param("id")
	response := pkg.NewResponse(ctx)

	if err := ctx.ShouldBind(&product); err != nil { // cek tipe data
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	_, err := govalidator.ValidateStruct(&product)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	// get file from request body
	file, header, _ := ctx.Request.FormFile("image")

	if file != nil {
		const maxFileSize = 5 * 1024 * 1024 // 5MB size limit
		if header.Size > maxFileSize {
			response.BadRequest("create data failed, upload file failed, file too large", "file size exceeds the 5MB limit")
			return
		}

		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			response.BadRequest("Create data failed, upload file failed", "Only jpg and png files are allowed")
			return
		}

		randomNumber := rand.Int()
		fileName := fmt.Sprintf("go-product-%d", randomNumber)
		uploadResult, err := h.UploadFile(ctx, file, fileName)
		if err != nil {
			response.BadRequest("create data failed, upload file failed", err.Error())
			return
		}

		picture := uploadResult.SecureURL
		product.Image = &picture
	}

	data, err := h.UpdateProduct(idParam, &product)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Product name already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	ctx.JSON(200, data)

}

func (h HandlerProduct) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := h.DeleteProductById(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	ctx.JSON(200, data)

}

func (h HandlerProduct) GetFavorite(ctx *gin.Context) {
	user := ctx.Param("userId")
	data, str, err := h.GetFavoritesProduct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, str)
		return
	}

	if len(*data) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No favorite found"})
		return
	}

	ctx.JSON(200, data)
}

func (h HandlerProduct) AddFavorite(ctx *gin.Context) {

	user := ctx.Param("userId")
	id := ctx.Param("productId")

	data, err := h.AddFavoriteProduct(user, id)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h HandlerProduct) DeleteFavorite(ctx *gin.Context) {
	user := ctx.Param("userId")
	id := ctx.Param("productId")

	data, err := h.DeleteFavoriteProduct(user, id)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}
