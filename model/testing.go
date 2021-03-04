package model

//TestComment used in tests
func TestComment() *Comment {
	return &Comment{
		ID:     0,
		PostID: 0,
		Name:   "ipsum",
		Email:  "someemail@mail.com",
		Body:   "test comment body",
	}
}

//TestUser used in tests
func TestUser() *User {
	return &User{
		ID:       0,
		Password: "password",
		Email:    "testemail@mail.com",
		Picture:   "",
		Name:     "TestUser",
	}
}

//TestPost used in tests
func TestPost() *Post {
	return &Post{
		ID:     0,
		UserID: 0,
		Title:  "Test Title",
		Body:   "Test post body",
	}
}
