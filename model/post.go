package model

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// Post struct
type Post struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	UserID int    `json:"userID"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

//NewPost ...
type NewPost struct {
	Title string `json:"title" example:"Tea"`
	Body  string `json:"body" example:"Tea is good for your health"`
}

//Validate post body and title
func (p *Post) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Body, validation.Required, validation.Length(1, 500)),
		validation.Field(&p.Title, validation.Required, validation.Length(1, 25)),
	)
}

//Sanitize sets id from JWT token and post_id to 0
func (p *Post) Sanitize(c echo.Context) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	JWTid := int(claims["id"].(float64))
	//sanitize
	p.ID = 0
	p.UserID = JWTid
}
