package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"net/http"

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

	data, err := h.GetAllProduct(search, sort, category)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerProduct) UpdateById(ctx *gin.Context) {
	product := models.Product{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&product); err != nil { // cek tipe data
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.UpdateProduct(id, &product)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, data)

}

func (h HandlerProduct) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := h.DeleteProductById(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, data)

}

func (h HandlerProduct) GetFavorite(ctx *gin.Context) {

	userId := ctx.Param("userId")
	data, err := h.GetFavoritesProduct(userId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, data)
}

func (h HandlerProduct) AddFavorite(ctx *gin.Context) {

	userId := ctx.Param("userId")
	productId := ctx.Param("productId")

	data, err := h.AddFavoriteProduct(userId, productId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h HandlerProduct) DeleteFavorite(ctx *gin.Context) {
	userId := ctx.Param("userId")
	productId := ctx.Param("productId")

	data, err := h.DeleteFavoriteProduct(userId, productId)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}
