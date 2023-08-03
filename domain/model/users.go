package model

import (
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

type ReadUsersWithPaginationReq struct {
	Fullname string
	Limit    int
	Offset   int
}

type ReadUsersWithPaginationResp struct {
	Users   []UserWithRole
	PerPage int
}

type UserWithRole struct {
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
	RoleName  string
}

type CountUsersReq struct {
	IsDeleted int
}

type CountUsersResp struct {
	Total int
}

type GenerateAccessTokenReq struct {
	UserUID    string
	RoleUID    string
	LastActive int
	ExpiredAt  int
}

type GenerateAccessTokenResp struct {
	Token string
}

type GenerateRefreshTokenReq struct {
	UserUID   string
	ExpiredAt int
}

type GenerateRefreshTokenResp struct {
	Token string
}

type ReadUserByEmailPasswordReq struct {
	Email    string
	Password string
	RoleID   int
}

type ReadUserByEmailPasswordResp struct {
	UserUID    string
	RoleUID    string
	LastActive time.Time
}
