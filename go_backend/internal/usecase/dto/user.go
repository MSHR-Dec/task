package dto

type SignUpInput struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignInInput struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SignInOutput struct {
	ShouldUpdatePassword bool `json:"should_update_password"`
}
