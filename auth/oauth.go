package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Monrevil/fonder/config"
	"github.com/Monrevil/fonder/model"

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

func handleCallback(db *gorm.DB, conf oauth2.Config, endpoint string) echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.FormValue("code")
		tok, err := conf.Exchange(c.Request().Context(), code)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		client := conf.Client(c.Request().Context(), tok)

		resp, err := client.Get(endpoint)
		if err != nil {
			return c.String(http.StatusBadRequest, "could not get info from google")
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
