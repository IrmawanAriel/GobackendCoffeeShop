package handlers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/models"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"
	"IrmawanAriel/goBackendCoffeeShop/pkg"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type HandlerUser struct {
	repositories.UserRepositoryInterface
}

func NewUser(r repositories.UserRepositoryInterface) *HandlerUser {
	return &HandlerUser{r}
}

func (h *HandlerUser) FetchById(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	user, err := h.GetUserById(id)
	if err != nil {
		response.NotFound("user not found", err.Error())
		return
	}

	response.Success("user fetched successfully", user)
}

func (h *HandlerUser) FetchAll(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	data, err := h.GetAllUser()
	if err != nil {
		response.InternalServerError("failed to fetch users", err.Error())
		return
	}

	response.Success("users fetched successfully", data)
}

func (h *HandlerUser) UpdateUserById(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	var data models.User

	if err := ctx.ShouldBind(&data); err != nil {
		response.BadRequest("invalid request data", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&data)
	if err != nil {
		response.BadRequest("validation failed", err.Error())
		return
	}

	res, err := h.UpdateUser(id, &data)
	if err != nil {
		response.NotFound("no such user", err.Error())
		return
	}

	response.Success("user updated successfully", res)
}

func (h *HandlerUser) Register(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	var data models.User

	if err := ctx.ShouldBind(&data); err != nil {
		response.BadRequest("invalid request data", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&data)
	if err != nil {
		response.BadRequest("validation failed", err.Error())
		return
	}

	data.Password, err = pkg.HashPassword(data.Password)
	if err != nil {
		response.InternalServerError("password hashing failed", err.Error())
		return
	}

	res, err := h.InsertUser(&data)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			response.BadRequest("account already exists", err.Error())
			return
		}
		response.InternalServerError("failed to register account", err.Error())
		return
	}

	response.Created("account registered successfully", res)
}

func (h *HandlerUser) DeleteUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	idUser := ctx.Param("id")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		response.BadRequest("invalid user ID", err.Error())
		return
	}

	res, err := h.DeleteUserById(id)
	if err != nil {
		response.NotFound("user not found", err.Error())
		return
	}

	response.Success("user deleted successfully", res)
}

func (r *HandlerUser) Login(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	var data models.Login

	if err := ctx.ShouldBind(&data); err != nil {
		response.BadRequest("invalid request data", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&data)
	if err != nil {
		response.BadRequest("validation failed", err.Error())
		return
	}

	result, err := r.GetByEmail(data.Email)
	if err != nil {
		response.NotFound("user not found", err.Error())
		return
	}

	if err := pkg.VerifyPassword(result.Password, data.Password); err != nil {
		response.Unauthorized("incorrect password", err.Error())
		return
	}

	jwt := pkg.NewJWT(result.Id, result.Email, result.Role)
	token, err := jwt.GenerateToken()
	if err != nil {
		response.InternalServerError("failed to generate token", err.Error())
		return
	}

	response.Success("login successful", token)
}

func (h *HandlerUser) Create(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	var data models.UserCreate

	if err := ctx.ShouldBind(&data); err != nil {
		response.BadRequest("invalid request data", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&data)
	if err != nil {
		response.BadRequest("validation failed", err.Error())
		return
	}

	data.Password, err = pkg.HashPassword(data.Password)
	if err != nil {
		response.InternalServerError("password hashing failed", err.Error())
		return
	}

	res, err := h.CreateUser(&data)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			response.BadRequest("account already exists", err.Error())
			return
		}
		response.InternalServerError("failed to create account", err.Error())
		return
	}

	response.Created("account created successfully", res)
}
