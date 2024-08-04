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
	data, err := h.GetAllProduct()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerProduct) FetchById(ctx *gin.Context) {
	product := models.Product{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&product); err != nil { // cek tipe data
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data, err := h.GetProductById(id, &product)
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
