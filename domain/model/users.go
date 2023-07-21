package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type User struct {
	ID        int
	UID       string
	FirstName string
	LastName  string
	Email     string
	Birthdate time.Time
	Username  string
	Password  string
	IsDeleted bool
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}

type TokenClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}
