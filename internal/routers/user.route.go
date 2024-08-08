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
	route.PATCH("/update/:id", handler.UpdateUserById)
	route.GET("/:id", handler.FetchById)
	route.DELETE("/delete/:id", handler.DeleteUser)
	route.POST("/create/", handler.Register)

}
