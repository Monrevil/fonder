package auth

import (
	"github.com/Monrevil/fonder/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

//GoogleCallback callback endpoint
//Google returns email+name
func GoogleCallback(db *gorm.DB) echo.HandlerFunc {
	conf := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       []string{config.Scope, "profile"},
		Endpoint:     google.Endpoint,
	}
	url := "https://www.googleapis.com/oauth2/v2/userinfo/"
	return handleCallback(db, conf, url)

}
