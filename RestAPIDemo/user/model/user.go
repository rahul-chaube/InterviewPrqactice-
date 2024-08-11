package model

import "github.com/google/uuid"

type User struct {
	Id    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Age   string `json:"age" bson:"age"`
	Email string `json:"email" bson:"email"`
}

func (u *User) SetUserId() {
	u.Id = uuid.NewString()
}
func (u *User) IsValidUser() bool {
	return u.Id == "" && u.Name == ""
}
