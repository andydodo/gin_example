package http

import (
	"github.com/gin-gonic/gin"
)

func (a *AppServer) privateRoutes(router *gin.RouterGroup) {
	router.GET("me", a.GetMeHandler)
	router.GET("users/:id", a.GetUserHandler)
	router.GET("link/:id", a.GetLinkHandler)
	router.POST("link", a.CreateLinkHandler)
	router.PUT("link/:id", a.UpdateLinkHandler)
	router.DELETE("link/:id", a.DeleteLinkHandler)
}

func (a *AppServer) publicRoutes(router *gin.RouterGroup) {
	router.POST("/register", a.RegisterUserHandler)
	router.POST("/login", a.LoginUserHandler)
	router.POST("/logout", a.LogoutUserHandler)
}
