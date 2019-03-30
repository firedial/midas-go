package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

    "github.com/firedial/midas-go/controller"
    "github.com/firedial/midas-go/model"
    "github.com/firedial/midas-go/config"

)

func main() {

    if config.IS_PRODUCTION {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.Default()
    r.Use(cors.Default())

    api := r.Group("/api/v1")
    {
        api.GET("/balance/", func(c *gin.Context) { c.JSON(200, controller.BalanceGet(c.Request.URL.Query())) } )
        api.POST("/balance/", func(c *gin.Context) { 
            var balance model.Balance
            c.BindJSON(&balance)
            c.JSON(200, controller.BalancePost(balance)) } )
        api.POST("/move/", func(c *gin.Context) { 
            var move model.Move
            c.BindJSON(&move)
            c.JSON(200, controller.MovePost(move)) } )
        api.GET("/kind/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("kind", c.Request.URL.Query())) } )
        api.GET("/purpose/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("purpose", c.Request.URL.Query())) } )
        api.GET("/place/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("place", c.Request.URL.Query())) } )
    }

    r.Run()
}
