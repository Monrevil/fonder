package auth

import (
	"github.com/Monrevil/fonder/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	f "golang.org/x/oauth2/facebook"
	"gorm.io/gorm"
)

//FacebookCallback redirect url
func FacebookCallback(db *gorm.DB) echo.HandlerFunc {
	conf := oauth2.Config{
		ClientID:     config.FacebookID,
		ClientSecret: config.FacebookSecret,
		Endpoint:     f.Endpoint,
		RedirectURL:  config.FacebookRedirect,
		Scopes:       []string{"email"},
	}
	url := "https://graph.facebook.com/v3.3/me/?fields=name,email"
	return handleCallback(db, conf, url)
}
