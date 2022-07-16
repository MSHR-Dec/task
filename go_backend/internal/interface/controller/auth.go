package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

const SessionKey = "task"

func setSession(ctx *gin.Context) error {
	session := sessions.Default(ctx)

	if isExistSession(session) {
		return nil
	}

	token := uuid.NewV4().String()
	session.Set(SessionKey, token)
	if err := session.Save(); err != nil {
		return oops.InternalServerError{Message: "failed to save session"}
	}
	return nil
}

func expireSession(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	if err := session.Save(); err != nil {
		return oops.InternalServerError{Message: "failed to save session"}
	}
	ctx.Set(SessionKey, nil)
	return nil
}

func isExistSession(s sessions.Session) bool {
	return s.Get(SessionKey) != nil
}