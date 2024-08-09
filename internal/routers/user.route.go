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

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/", middleware.AuthJwtMiddleware("admin"), handler.FetchAll)
	route.POST("/login", handler.Login)
	route.GET("/:id", middleware.AuthJwtMiddleware("user"), handler.FetchById)
	route.POST("/create/", middleware.AuthJwtMiddleware("admin"), handler.Register)
	route.POST("/register", handler.Register)
	route.DELETE("/delete/:id", middleware.AuthJwtMiddleware("admin"), handler.DeleteUser)
	route.PATCH("/update/:id", middleware.AuthJwtMiddleware("user"), handler.UpdateUserById)

}
