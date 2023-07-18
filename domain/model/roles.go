package model

import "time"

type Role struct {
	ID        int
	UID       string
	Name      string
	IsDeleted bool
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}

type ReadRolesByIDReq struct {
	ID int
}
