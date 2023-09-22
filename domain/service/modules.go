package service

import (
	"encoding/json"
	"github/yogabagas/join-app/shared/util"
	"net/http"
)

type CreateModulesReq struct {
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	File           string           `json:"-"`
	ModuleMaterial []ModuleMaterial `json:"module_materials"`
}

type ModuleMaterial struct {
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

func (c CreateModulesReq) SetCreateModuleReq(r *http.Request) (CreateModulesReq, error) {
	c.Name = r.FormValue("name")
	c.Description = r.FormValue("description")
	c.ModuleMaterial = parseRequestModuleMaterial(r.FormValue("module_materials"))

	filename, err := parseFile(r)
	c.File = filename

	return c, err
}

func parseFile(r *http.Request) (string, error) {
	filename, err := util.ParseFileUpload(r, "file", "storage")

	return filename, err
}

func parseRequestModuleMaterial(request string) (materials []ModuleMaterial) {
	json.Unmarshal([]byte(request), &materials)

	return materials
}
