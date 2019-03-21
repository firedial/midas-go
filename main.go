package main

import (
    "github.com/gin-gonic/gin"

    //"github.com/firedial/midas/db"
    "github.com/firedial/go-midas/controller"
    "github.com/firedial/go-midas/model"

)

func main() {
    r := gin.Default()

    balance := r.Group("/api/v1/balance")
    {
        balance.GET("/", func(c *gin.Context) { c.JSON(200, controller.BalanceGet(c.Request.URL.Query())) } )
        // balance.POST("/", func(c *gin.Context) { c.JSON(200, controller.BalancePost(c.GetPostFormMap())) } )
        balance.POST("/", func(c *gin.Context) { 
            var balance model.Balance
            c.BindJSON(&balance)
            c.JSON(200, controller.BalancePost(balance)) } )
    }

    r.GET("/api/v1/kind/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("kind", c.Request.URL.Query())) } )
    r.GET("/api/v1/purpose/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("purpose", c.Request.URL.Query())) } )
    r.GET("/api/v1/place/", func(c *gin.Context) { c.JSON(200, controller.AttributeGet("place", c.Request.URL.Query())) } )
    //r.GET("/", func(c *gin.Context) {
    //    c.String(200, "Hello, World")
    //})
    //db.Init()
    r.Run()
}
