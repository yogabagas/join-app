package service

type CreateUsersReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
	RoleID    int    `json:"role_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedBy string `json:"-"`
}
