package handler

import (
	"encoding/json"
	"mnt/c/Users/DELL/nix/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetComment(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/comments/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")
	db := model.InitDB()
	echoHadler := GetComment(db)
	// Assertions
	if assert.NoError(t, echoHadler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	//	assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestCreateComment(t *testing.T) {
	// Setup
	e := echo.New()
	db := model.InitDB()
	user := model.TestUser()
	token := user.GetJWTRaw()

	testCases := []struct {
		name    string
		c       func() *model.Comment
		expCode int
	}{
		{
			name: "Valid",
			c: func() *model.Comment {
				return model.TestComment()
			},
			expCode: http.StatusCreated,
		},
		{
			name: "No body",
			c: func() *model.Comment {
				comment := model.TestComment()
				comment.Body = ""
				return comment
			},
			expCode: http.StatusBadRequest,
		},
		{
			name: "No name",
			c: func() *model.Comment {
				comment := model.TestComment()
				comment.Name = ""
				return comment
			},
			expCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		byte, err := json.Marshal(tc.c())
		if err != nil {
			t.Fatal(err)
		}
		commentJSON := string(byte)
	
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(commentJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		//Set jwt into context, will be handled by JWT middleware
		//Without jwt user in unauthorized
		ctx.Set("user", token)
		echoHandler := SaveComment(db)
		t.Run(tc.name, func(t *testing.T) {
			// Assertions
			if assert.NoError(t, echoHandler(ctx)) {
				assert.Equal(t, tc.expCode, rec.Code)
				//assert.Equal(t, userJSON, rec.Body.String())
			}
		})
	}
}
