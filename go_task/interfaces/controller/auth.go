package controller

import (
	"fmt"
	"github.com/MSHR-Dec/MSHR-Doc/mypkg/oops"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const SessionKey = "task"

func setSession(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	//session.Options(sessions.Options{
	//	HttpOnly: true,
	//	//SameSite: http.SameSiteLaxMode,
	//})

	if isExistSession(session) {
		return nil
	}

	token := uuid.NewString()
	fmt.Println(token)
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
