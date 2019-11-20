package auth

import "net/http"

var auths []Auth

type Auth struct {
	Username   string
	Password   string
	Token      string
	Expiration string
}

type App struct {
	Client *http.Client
}

type AuthenticityToken struct {
	Token string
}
