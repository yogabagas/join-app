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

type LoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	Account     LoginUsersRes `json:"account"`
	AccessToken string        `json:"access_token"`
}

type LoginUsersRes struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Birthdate *string `json:"birthdate,omitempty"`
	Email     *string `json:"email,omitempty"`
	RoleID    *int    `json:"role_id,omitempty"`
	Username  *string `json:"username,omitempty"`
}
