package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerUser struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) FetchById(ctx *gin.Context) {
	user := models.User{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&user); err != nil { // cek tipe data, jika benar assign
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := h.GetUserById(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, user)

}
