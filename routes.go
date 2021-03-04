package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"mnt/c/Users/DELL/nix/config"
	_ "mnt/c/Users/DELL/nix/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	h "mnt/c/Users/DELL/nix/handler"
)

func setRoutes(e *echo.Echo, db *gorm.DB) {
	isLoggedIn := middleware.JWT(config.JWTSecret)
	// Routes
	comment := e.Group("/comment/")
	comment.POST("", h.SaveComment(db), isLoggedIn)
	comment.GET("", h.ListComments(db))
	comment.GET(":id", h.GetComment(db))
	comment.PUT("", h.UpdateComment(db))
	comment.DELETE(":id", h.DeleteComment(db))

	p := e.Group("/post/")
	p.POST("", h.SavePost(db), isLoggedIn)
	p.GET("", h.ListPosts(db))
	p.GET(":id", h.GetPost(db))
	p.PUT("", h.UpdatePost(db))
	p.DELETE(":id", h.DeletePost(db))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/home/", home)
	e.GET("/googleCallback/", googleCallback(db))
	e.POST("/signup/", signup(db))
	e.GET("/login/", login(db))
}
