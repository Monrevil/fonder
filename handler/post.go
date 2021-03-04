package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/Monrevil/fonder/model"
)

// ListPosts godoc
// @Summary List posts
// @Description get list of posts posts
// @Tags posts
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Post
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /post/ [get]
func ListPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := model.PingDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error()+" Could not establish db connection")
		}
		var posts []model.Post
		result := db.Find(&posts)

		if result.Error != nil {
			return c.String(http.StatusInternalServerError, result.Error.Error()+" Could not get data from db")
		}
		return c.JSON(http.StatusOK, posts)
	}
}

// SavePost godoc
// @Summary Add a post
// @Description add by json post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body model.NewPost true "Add post"
// @Security ApiKeyAuth
// @Success 201 {object} model.Post
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /post/ [post]
func SavePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := new(model.Post)
		if err := c.Bind(post); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := post.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		//get info from jwt
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		JWTid := int(claims["id"].(float64))
		//sanitize
		sanitized := &model.Post{
			ID:     0,
			UserID: JWTid,
			Title:  post.Title,
			Body:   post.Body,
		}
		model.PingDB(db)
		db.Create(sanitized)

		return c.JSON(http.StatusCreated, sanitized)
	}
}

// GetPost godoc
// @Summary Show a post
// @Description get post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Post
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /post/{id} [get]
func GetPost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := model.PingDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		post := &model.Post{}
		result := db.Take(post, id)

		if result.Error != nil {
			return c.String(http.StatusNotFound, result.Error.Error())
		}
		return c.JSON(http.StatusOK, post)
	}
}

// UpdatePost godoc
// @Summary Update a post
// @Description Update by json post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param  post body model.Post true "Update post"
// @Success 200 {object} model.Post
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /post/ [put]
func UpdatePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := &model.Post{}
		if err := c.Bind(post); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if err := db.Take(&model.Post{}, post.ID); err != nil {
			return c.JSON(http.StatusNotFound, "record not found")
		}
		db.Save(post)
		return c.JSON(http.StatusOK, post)
	}
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete by post ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Param  id path int true "post ID" Format(int64)
// @Success 204 {object} model.Post
// @Router /post/{id} [delete]
func DeletePost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&model.Post{}, id)
		return c.JSON(http.StatusNoContent, map[string]string{id: "Was deleted"})
	}
}
