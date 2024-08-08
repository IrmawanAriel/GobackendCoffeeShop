package routers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/handlers"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(d)
	handler := handlers.NewUser(repo)

	route.GET("/", handler.FetchAll)
	route.POST("/login", handler.Login)
	route.GET("/:id", handler.FetchById)
	route.POST("/create/", handler.Register)
	route.POST("/register", handler.Register)
	route.DELETE("/delete/:id", handler.DeleteUser)
	route.PATCH("/update/:id", handler.UpdateUserById)

}
