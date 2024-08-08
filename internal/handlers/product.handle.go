package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerProduct struct {
	*repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *HandlerProduct {
	return &HandlerProduct{r}
}

func (h *HandlerProduct) PostProduct(ctx *gin.Context) {
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil { // cek tipe data
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respone, err := h.CreateProduct(&product) // create the product
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	data, err := h.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No such product"})
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerProduct) UpdateById(ctx *gin.Context) {
	product := models.Product{}
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	if err := ctx.ShouldBind(&product); err != nil { // cek tipe data
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid favorite ID"})
		return
	}

	data, err := h.UpdateProduct(id, &product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "data is invalid"})
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
	userId, err2 := strconv.Atoi(user)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	data, str, err := h.GetFavoritesProduct(userId)
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

	userId, err2 := strconv.Atoi(user)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	data, err := h.AddFavoriteProduct(userId, productId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h HandlerProduct) DeleteFavorite(ctx *gin.Context) {
	user := ctx.Param("userId")
	id := ctx.Param("productId")

	userId, err2 := strconv.Atoi(user)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}
	productId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product ID"})
		return
	}

	data, err := h.DeleteFavoriteProduct(userId, productId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}
