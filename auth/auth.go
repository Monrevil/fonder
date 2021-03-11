package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	f "golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"

	"gorm.io/gorm"

	"github.com/Monrevil/fonder/config"
	"github.com/Monrevil/fonder/model"
)

// Home godoc
// @Summary Log in with OAUTH2
// @Description Use provided link to log in with Google/Facebook
// @Tags auth
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Comment
// @Router /home/ [get]
func Home(c echo.Context) error {
	conf := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       []string{config.Scope, "profile"},
		Endpoint:     google.Endpoint,
	}
	url := conf.AuthCodeURL(config.State)
	facebook := oauth2.Config{
		ClientID:     config.FacebookID,
		ClientSecret: config.FacebookSecret,
		//x/oauth2 uses outdated enpoint for Facebook
		Endpoint:    f.Endpoint,
		RedirectURL: config.FacebookRedirect,
		Scopes:      []string{"email"},
	}
	facebookURL := facebook.AuthCodeURL(config.State)
	twitter := oauth2.Config{
		ClientID:     config.TwitterID,
		ClientSecret: config.TwitterSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.twitter.com/oauth/authorize",
			TokenURL: "https://api.twitter.com/oauth/access_token",
		},
		RedirectURL: config.TwitterRedirect,
		Scopes:      []string{"email"},
	}
	twitterURL := twitter.AuthCodeURL(config.State)
	return c.JSON(200, map[string]string{
		"Google login: ":       url,
		"Facebook login: ":     facebookURL,
		"Login with Twitter: ": twitterURL,
	})
}

// Signup godoc
// @Summary Signup with app
// @Description Register: email should be unique, pasword from 6 to 25 char long
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  user body model.User true "Register user"
// @Success 200 {object} model.User
// @Failure 400
// @Router /signup/ [post]
func Signup(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &model.User{}
		c.Bind(user)
		err := user.Validate()
		if err != nil {
			return c.JSON(400, err)
		}
		user.ID = 0
		err = db.Where("email = ?", user.Email).First(user).Error
		if err == nil {
			return c.String(http.StatusBadRequest, "User with this email already exists")
		}
		db.Save(user)
		return c.JSON(http.StatusOK, user)
	}
}

// Login godoc
// @Summary Login with the app
// @Description Login with app if you have registered with us
// @Tags auth
// @Accept  json
// @Produce  json
// @Param  name query string true "name" Format(string)
// @Param  password query string true "password" Format(string)
// @Success 200 {object} model.User
// @Failure 400
// @Failure 401 
// @Router /login/ [get]
func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := &model.User{}
		user.Name = c.QueryParam("name")
		user.Password = c.QueryParam("password")

		if len(user.Name) == 0 || len(user.Password) < 6 {
			return c.String(http.StatusBadRequest, "No username or short password")
		}

		err := db.Where("name = ? AND password = ?", user.Name, user.Password).First(user).Error
		if err != nil {
			return c.String(http.StatusUnauthorized, "No such user+passwrod")
		}
		jwt, err := user.GetJWT()
		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.String(http.StatusOK, jwt)
	}
}
