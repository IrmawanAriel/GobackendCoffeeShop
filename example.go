package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func exam() {
	router := gin.Default()

	router.GET("/test", example)
	router.GET("/query", queryString)
	router.GET("/param/:hoby", paramString)
	router.POST("/body", reqBody)

	router.Run(":8081")

}

func example(ctx *gin.Context) {
	ctx.String(200, "hello worlds")
}

// ! http://localhost:8081/query?page=2&limit=10
func queryString(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	ctx.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
	})
}

// ! http://localhost:8081/params/makan
func paramString(ctx *gin.Context) {
	hoby := ctx.Param("hoby")

	ctx.JSON(200, gin.H{
		"hoby": hoby,
	})
}

type User struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func reqBody(ctx *gin.Context) {
	var data User

	if err := ctx.ShouldBind(&data); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, data)
}
