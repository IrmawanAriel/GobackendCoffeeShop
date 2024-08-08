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

	route.GET("/", handler.FetchAll) // sorting and seacrh are included
	route.GET("/:id", handler.FetchById)
	route.POST("/", handler.PostProduct)
	route.PATCH("/update/:id", handler.UpdateById)
	route.DELETE("/delete/:id", handler.DeleteProduct)

	// favorite
	route.GET("/favorite/:userId/", handler.GetFavorite)
	// route.GET("/favorite/:userId/", handler.SearchParams)
	route.POST("/favorite/add/:userId/:productId", handler.AddFavorite)
	route.DELETE("/favorite/delete/:userId/:productId", handler.DeleteFavorite)

}
