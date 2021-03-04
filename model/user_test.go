package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *User
		isValid bool
	}{
		{
			name: "Valid",
			u: func() *User {
				return TestUser()
			},
			isValid: true,
		},
		{
			name: "Empy Password",
			u: func() *User {
				u := TestUser()
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Empty Email",
			u: func() *User {
				u := TestUser()
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "Invalid Email",
			u: func() *User {
				u := TestUser()
				u.Email = "ivalid"
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
