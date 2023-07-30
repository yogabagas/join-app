package model

import "time"

type Authz struct {
	ID         int
	UID        string
	UserUID    string
	RoleUID    string
	IsDeleted  bool
	LastActive time.Time
	CreatedBy  string
	CreatedAt  time.Time
	UpdatedBy  string
	UpdatedAt  time.Time
}
