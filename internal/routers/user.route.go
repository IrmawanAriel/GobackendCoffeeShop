package routers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/handlers"
	"IrmawanAriel/goBackendCoffeeShop/internal/middleware"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	var repo repositories.UserRepositoryInterface = repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/", middleware.AuthJwtMiddleware("admin"), handler.FetchAll)
	route.POST("/login", handler.Login)
	route.GET("/:id", middleware.AuthJwtMiddleware(""), handler.FetchById)
	route.POST("/register", handler.Register)
	route.POST("/create", middleware.AuthJwtMiddleware("admin"), handler.Create)
	route.DELETE("/delete/:id", middleware.AuthJwtMiddleware("admin"), handler.DeleteUser)
	route.PATCH("/update/:id", middleware.AuthJwtMiddleware(""), handler.UpdateUserById)

}
