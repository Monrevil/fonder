package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"

	"github.com/Monrevil/fonder/config"
)

// User contains info: email, id, name
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Picture   string `json:"picture"`
	Name     string `json:"name"`
}

//Validate user
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u, 
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 25)),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 25)),
	)
}


// GetJWTRaw returns unsigned jwt token
func (u *User) GetJWTRaw() *jwt.Token {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["name"] = u.Name
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token
}
//GetJWT retruns signed string jwt for user
//With prefix e.g.: `Bearer <token>`
func (u *User) GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["name"] = u.Name
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token.
	t, err := token.SignedString([]byte(config.JWTSecret))
	t = "Bearer " + t
	if err != nil {
		return "", err
	}
	return t, nil
}

//InDB gets user Data by email, if exists in DB
//or create user if there is no such email
//Used in order to get an id for users registered with social media
// *overwrites info excep email
func (u *User) InDB(db *gorm.DB) bool {
	err := db.Where("email = ?", u.Email).First(u).Error
	if err != nil {
		db.Save(u)
		return false
	}
	return true
}
