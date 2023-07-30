package service

type CreateResourcesReq struct {
	Name      string `json:"name"`
	Type      int    `json:"type"`
	Action    string `json:"action"`
	ParentUID string `json:"parent_uid"`
	CreatedBy string `json:"-"`
}
