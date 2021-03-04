package handler

import (
	"encoding/json"
	"github.com/Monrevil/fonder/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	// Setup
	e := echo.New()
	db := model.InitDB()
	user := model.TestUser()
	token := user.GetJWTRaw()

	testCases := []struct {
		name    string
		c       func() *model.Post
		expCode int
	}{
		{
			name: "Valid",
			c: func() *model.Post {
				return model.TestPost()
			},
			expCode: http.StatusCreated,
		},
		{
			name: "No post Body",
			c: func() *model.Post {
				post := model.TestPost()
				post.Body = ""
				return post
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "No post Title",
			c: func() *model.Post {
				post := model.TestPost()
				post.Title = ""
				return post
			},
			expCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		byte, err := json.Marshal(tc.c())
		if err != nil {
			t.Fatal(err)
		}
		postJSON := string(byte)

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		//Set jwt into context, will be handled by JWT middleware
		//Without jwt user in unauthorized
		ctx.Set("user", token)
		echoHandler := SavePost(db)
		t.Run(tc.name, func(t *testing.T) {
			// Assertions
			if assert.NoError(t, echoHandler(ctx)) {
				assert.Equal(t, tc.expCode, rec.Code)
				//assert.Equal(t, userJSON, rec.Body.String())
			}
		})
	}
}
