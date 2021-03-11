package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Monrevil/fonder/config"
	"github.com/Monrevil/fonder/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	f "golang.org/x/oauth2/facebook"
	"gorm.io/gorm"
)

//FacebookCallback redirect url
func FacebookCallback(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		conf := oauth2.Config{
			ClientID:     config.FacebookID,
			ClientSecret: config.FacebookSecret,
			//x/oauth2 uses outdated enpoint for Facebook
			Endpoint:    f.Endpoint,
			RedirectURL: config.FacebookRedirect,
			Scopes:      []string{"email"},
		}
		code := c.FormValue("code")
		tok, err := conf.Exchange(c.Request().Context(), code)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		client := conf.Client(c.Request().Context(), tok)
		url := "https://graph.facebook.com/v3.3/me/?fields=name,email"
		resp, err := client.Get(url)
		if err != nil {
			return c.String(http.StatusBadRequest, "could not get info from facebook")
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		user := new(model.User)
		json.Unmarshal(bytes, &user)

		user.InDB(db)
		token, err := user.GetJWT()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, token)
	}

}
