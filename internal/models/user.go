package models

import "github.com/gorilla/sessions"

type User struct {
	Email     string
	FirstName string
	LastName  string
	UserID    string
	AvatarURL string
	Location  string
}

func NewUserFromSession(userSession *sessions.Session) User {
	return User{
		Email:     userSession.Values["email"].(string),
		FirstName: userSession.Values["first-name"].(string),
		LastName:  userSession.Values["last-name"].(string),
		UserID:    userSession.Values["user-id"].(string),
		AvatarURL: userSession.Values["avatar"].(string),
		Location:  userSession.Values["location"].(string),
	}
}
