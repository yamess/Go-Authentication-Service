package models

type Login struct {
	Email    string
	Password string
}

type Token struct {
	Token        string
	RefreshToken string
}

type LogOut struct {
	Token
}
