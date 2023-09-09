package service

type CreateModulesReq struct {
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	File           string           `json:"file"`
	ModuleMaterial []ModuleMaterial `json:"module_materials"`
}

type ModuleMaterial struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}
