package model

import validation "github.com/go-ozzo/ozzo-validation"

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
