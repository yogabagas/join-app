package model

import "time"

type Resource struct {
	ID        int
	UID       string
	Name      string
	ParentUID string
	Type      int
	Action    string
	IsDeleted bool
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
}
