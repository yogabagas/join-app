package controller

type AppController struct {
	AuthzController     interface{ AuthzController }
	UsersController     interface{ UsersController }
	ResourcesController interface{ ResourcesController }
	RolesController     interface{ RolesController }
}
