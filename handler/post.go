package handler

import (
	"net/http"

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
// @Produce xml
// @Success 200 {array} model.Post
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /post/ [get]
func ListPosts(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var posts []model.Post
		if err := listAll(db, &posts); err != nil {
			return respond(c, http.StatusInternalServerError, err)
		}
		return respond(c, http.StatusOK, posts)
	}
}

// SavePost godoc
// @Summary Add a post
// @Description add by json post
// @Tags posts
// @Accept  json
// @Produce  json
// @Produce xml
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
		if code, err := saveBody(c, db, post); err != nil {
			return respond(c, code, err.Error())
		}
		return respond(c, http.StatusOK, post)

	}

}



// GetPost godoc
// @Summary Show a post
// @Description get post by ID
// @Tags posts
// @Accept  json
// @Produce  json
// @Produce xml
// @Param id path int true "Post ID"
// @Success 200 {object} model.Post
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /post/{id} [get]
func GetPost(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		post := &model.Post{}
		if code, err := getByID(c, db, post); err != nil {
			return respond(c, code, err.Error())
		}
		return respond(c, http.StatusOK, post)
	}

}

// UpdatePost godoc
// @Summary Update a post
// @Description Update by json post
// @Tags posts
// @Accept  json
// @Produce  json
// @Produce xml
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
			return respond(c, http.StatusBadRequest, err.Error())
		}
		if err := db.Take(&model.Post{}, post.ID); err != nil {
			return respond(c, http.StatusNotFound, "record not found")
		}
		db.Save(post)

		return respond(c, http.StatusOK, post)
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
		return respond(c, http.StatusNoContent, map[string]string{id: "Was deleted"})
	}
}
