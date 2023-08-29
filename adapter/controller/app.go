package controller

type AppController struct {
	AccessController    interface{ AccessController }
	AuthzController     interface{ AuthzController }
	UsersController     interface{ UsersController }
	ResourcesController interface{ ResourcesController }
	RolesController     interface{ RolesController }
}
