package model

import "time"

type Module struct {
	ID          int
	UID         string
	Name        string
	Description string
	File        string
	IsDeleted   bool
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedBy   string
	UpdatedAt   time.Time
}

type ModuleMaterial struct {
	ID          int
	UID         string
	ModuleUID   string
	Topic       string
	Description string
	IsDeleted   bool
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedBy   string
	UpdatedAt   time.Time
}
