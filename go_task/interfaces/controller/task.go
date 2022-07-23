package controller

import (
	"net/http"
	"strconv"

	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
	"github.com/gin-gonic/gin"

	"github.com/MSHR-Dec/task/go_task/application"
	"github.com/MSHR-Dec/task/go_task/application/dto"
)

type TaskController struct {
	taskInteractor application.TaskInteractor
}

func NewTaskController(taskInteractor application.TaskInteractor) TaskController {
	return TaskController{
		taskInteractor: taskInteractor,
	}
}

func (c TaskController) Add(ctx *gin.Context) {
	var input dto.TaskAddInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid url parameter"})
		return
	}
	input.UserID = userID

	output, err := c.taskInteractor.Add(input)
	if err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func (c TaskController) Edit(ctx *gin.Context) {
	var input dto.TaskEditInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid request body"})
		return
	}

	if err := c.taskInteractor.Update(input); err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "edit task successfully"})
}

func (c TaskController) ListByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		oops.ResponseError(ctx, oops.BadRequest{Message: "invalid url parameter"})
		return
	}

	input := dto.TaskListInput{
		UserID: userID,
	}

	output, err := c.taskInteractor.ListByUserID(input)
	if err != nil {
		oops.ResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}
