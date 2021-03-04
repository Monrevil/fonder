package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/Monrevil/fonder/config"
	_ "github.com/Monrevil/fonder/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	a "github.com/Monrevil/fonder/auth"
	h "github.com/Monrevil/fonder/handler"
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

	e.GET("/home/", a.Home)
	e.GET("/googleCallback/", a.GoogleCallback(db))
	e.POST("/signup/", a.Signup(db))
	e.GET("/login/", a.Login(db))
}
