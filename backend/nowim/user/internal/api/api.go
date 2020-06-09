package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nowim.user/internal/domain"
)

func SetUpHandlers(e *gin.Engine) {
	userService := NewUserService(domain.NewPostgresUserRepo())
	e.Handle(http.MethodPost, "/api/user/create", userService.CreateUser)
	e.Handle(http.MethodPost, "/api/user/login", userService.Login)
	e.Handle(http.MethodPost, "/api/user/fetch", userService.FetchUser)
}
