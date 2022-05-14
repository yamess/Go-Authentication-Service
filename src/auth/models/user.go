package models

import "github.com/gofrs/uuid"

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserRegister struct {
	User
	VerifyPassword string
}
