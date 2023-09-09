package controller

type AppController struct {
	UsersController     interface{ UsersController }
	ResourcesController interface{ ResourcesController }
	RolesController     interface{ RolesController }
	CoursesController   interface{ ModulesController }
}
