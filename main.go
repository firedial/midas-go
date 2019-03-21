package main

import (
    "github.com/gin-gonic/gin"

    "github.com/firedial/go-midas/controller"
    "github.com/firedial/go-midas/model"

)

func main() {
    r := gin.Default()

    api := r.Group("/api/v1")
    {
        api.GET("/balance/", func(c *gin.Context) { c.JSON(200, controller.BalanceGet(c.Request.URL.Query())) } )
        api.POST("/balance/", func(c *gin.Context) { 
            var balance model.Balance
            c.BindJSON(&balance)
            c.JSON(200, controller.BalancePost(balance)) } )
        api.GET("/kind/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("kind", c.Request.URL.Query())) } )
        api.GET("/purpose/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("purpose", c.Request.URL.Query())) } )
        api.GET("/place/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("place", c.Request.URL.Query())) } )
    }

    r.Run()
}
