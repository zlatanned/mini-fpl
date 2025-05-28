package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/zlatanned/mini-fpl/internal/auth"
    "github.com/zlatanned/mini-fpl/configs"
)

func main() {
    configs.LoadEnv()
    r := gin.Default()

    // Health check
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    });

    r.POST("/auth/register", auth.Register);

    r.POST("/auth/login", auth.Login);

    r.Run(":3061") // listen and serve on 0.0.0.0:3061
}