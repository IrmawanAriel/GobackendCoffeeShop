package routers

import (
	// "biFebriansyah/goback/internal/handlers"
	// "biFebriansyah/goback/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func movie(g *gin.Engine, d *sqlx.DB) {
	// route := g.Group("/movie")

	// repo := repository.NewMovie(d)
	// handler := handlers.NewMovie(repo)

	// route.GET("/", handler.FetchAll)
	// route.POST("/", handler.PostMovie)
}
