package service

import (
	"encoding/json"
)

type CreateModulesReq struct {
	UID            string           `json:"uid"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	File           string           `json:"-"`
	ModuleMaterial []ModuleMaterial `json:"module_materials"`
}

type UpdateModulesReq struct {
	UID            string           `json:"uid"`
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	File           string           `json:"file"`
	NewFile        string           `json:"-"`
	ModuleMaterial []ModuleMaterial `json:"module_materials"`
}

type ModuleMaterial struct {
	UID         string `json:"uid"`
	ModuleUID   string `json:"module_uid"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
}

type GetModulesWithPaginationReq struct {
	UID   string `json:"uid"`
	Name  string `json:"name"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type ModuleResp struct {
	UID             string               `json:"uid"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	File            string               `json:"file"`
	ModuleMaterials []ModuleMaterialResp `json:"module_materials"`
}

type ModuleMaterialResp struct {
	UID         string `json:"uid"`
	ModuleUID   string `json:"module_uid"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
}

type GetModulesWithPaginationResp struct {
	Modules    []ModuleResp `json:"modules"`
	Pagination Pagination   `json:"pagination"`
}

func ParseRequestModuleMaterial(request string) (materials []ModuleMaterial) {
	json.Unmarshal([]byte(request), &materials)

	return materials
}
