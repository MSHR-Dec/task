package controller

import (
	"net/http"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
	"github.com/gin-gonic/gin"

	"github.com/MSHR-Dec/task/go_task/application"
	"github.com/MSHR-Dec/task/go_task/application/dto"
)

type UserController struct {
	userInteractor application.UserInteractor
}

func NewUserController(userInteractor application.UserInteractor) UserController {
	return UserController{
		userInteractor: userInteractor,
	}
}

func (c UserController) SignUp(ctx *gin.Context) {
	var input dto.SignUpInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	output, err := c.userInteractor.SignUp(input)
	if err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	if err := setSession(ctx); err != nil {
		oops.ResponseError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, output)
}

func (c UserController) SignIn(ctx *gin.Context) {
	var input dto.SignInInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	output, err := c.userInteractor.SignIn(input)
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

func (c UserController) EditProfile(ctx *gin.Context) {
	var input dto.EditProfileInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	if err := c.userInteractor.EditProfile(input); err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "edit profile successfully"})
}
