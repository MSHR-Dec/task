package dto

type SignUpInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpOutput struct {
	ID int `json:"id"`
}

type SignInInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInOutput struct {
	ID                   int  `json:"id"`
	ShouldUpdatePassword bool `json:"should_update_password"`
}

type EditProfileInput struct {
	ID       int    `json:"id" binding:"required"`
	Password string `json:"password"`
}
