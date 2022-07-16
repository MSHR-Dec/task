package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MSHR-Dec/task/go_backend/internal/usecase/dto"
	"github.com/MSHR-Dec/task/go_backend/internal/usecase/input"
	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

type UserController struct {
	userInteractor input.UserInputPort
}

func NewUserController(userInteractor input.UserInputPort) UserController {
	return UserController{
		userInteractor: userInteractor,
	}
}

func (c UserController) SignUp(ctx *gin.Context) {
	var signUpDTO dto.SignUpInput
	if err := ctx.ShouldBindJSON(&signUpDTO); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	if err := c.userInteractor.SignUp(signUpDTO); err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	if err := setSession(ctx); err != nil {
		oops.ResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, "sign up successfully")
}

func (c UserController) SignIn(ctx *gin.Context) {
	var signInDTO dto.SignInInput
	if err := ctx.ShouldBindJSON(&signInDTO); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	output, err := c.userInteractor.SignIn(signInDTO)
	if err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	if err = setSession(ctx); err != nil {
		oops.ResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func (c UserController) SignOut(ctx *gin.Context) {
	if err := expireSession(ctx); err != nil {
		oops.ResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "sign out successfully"})
}
