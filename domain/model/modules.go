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

type ReadModulesWithPaginationReq struct {
	UID    string
	Name   string
	Limit  int
	Offset int
}

type ReadModulesWithPaginationResp struct {
	Modules []Module
	PerPage int
}

type CountModulesReq struct {
	IsDeleted int
}

type CountModulesResp struct {
	Total int
}
