package model

import "time"

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
