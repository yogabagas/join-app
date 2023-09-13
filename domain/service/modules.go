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

func (c CreateModulesReq) SetCreateModuleReq(r *http.Request) (CreateModulesReq, error) {
	c.Name = r.FormValue("name")
	c.Description = r.FormValue("description")
	c.ModuleMaterial = parseRequestModuleMaterial(r.FormValue("module_materials"))

	filename, err := parseFile(r)
	c.File = filename

	return c, err
}

type ModuleMaterial struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}

func parseFile(r *http.Request) (string, error) {
	filename, err := util.ParseFileUpload(r, "file", "storage")

	return filename, err
}

func parseRequestModuleMaterial(request string) (materials []ModuleMaterial) {
	json.Unmarshal([]byte(request), &materials)

	return materials
}
