package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
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

func (h *HandlerUser) FetchAll(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, data)
}

func (h *HandlerUser) UpdateUserById(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	var data models.User

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := govalidator.ValidateStruct(&data)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	res, err := h.UpdateUser(id, &data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No such a user"})
		return
	}
	ctx.JSON(200, res)
}

func (h *HandlerUser) Register(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	var data models.User

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	_, err := govalidator.ValidateStruct(&data)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	res, err := h.InsertUser(&data)

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, res)

}

func (h HandlerUser) DeleteUser(ctx *gin.Context) {
	idUser := ctx.Param("id")
	id, _ := strconv.Atoi(idUser)

	res, err := h.DeleteUserById(id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, res)

}
