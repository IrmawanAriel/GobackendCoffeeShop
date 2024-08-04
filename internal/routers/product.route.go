package routers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/handlers"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.GET("/", handler.FetchAll)
	route.POST("/", handler.PostProduct)
	route.PATCH("/update/:id", handler.FetchById)

}
