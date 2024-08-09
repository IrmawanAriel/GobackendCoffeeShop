package routers

import (
	"IrmawanAriel/goBackendCoffeeShop/internal/handlers"
	"IrmawanAriel/goBackendCoffeeShop/internal/middleware"
	"IrmawanAriel/goBackendCoffeeShop/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(d)
	handler := handlers.NewProduct(repo)

	route.GET("/", handler.FetchAll) // sorting and seacrh are included
	route.GET("/:id", middleware.AuthJwtMiddleware("user"), handler.FetchById)
	route.POST("/", middleware.AuthJwtMiddleware("admin"), handler.PostProduct)
	route.PATCH("/update/:id", middleware.AuthJwtMiddleware("admin"), handler.UpdateById)
	route.DELETE("/delete/:id", middleware.AuthJwtMiddleware("admin"), handler.DeleteProduct)

	// favorite
	route.GET("/favorite/:userId/", middleware.AuthJwtMiddleware("user"), handler.GetFavorite)
	// route.GET("/favorite/:userId/", handler.SearchParams)
	route.POST("/favorite/add/:userId/:productId", middleware.AuthJwtMiddleware("user"), handler.AddFavorite)
	route.DELETE("/favorite/delete/:userId/:productId", middleware.AuthJwtMiddleware("user"), handler.DeleteFavorite)

}
