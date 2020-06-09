package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"nowim.user/internal/auth"
	"nowim.user/internal/domain"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserReply struct {
	UserID int64 `json:"userID"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReply struct {
	Token string `json:"token"`
}

type FetchUserRequest struct {
}

type FetchUserReply struct {
	Users map[int64]*domain.User `json:"users"`
}

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s *UserService) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		log.Infof("invalid request, err: %+v", err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	user := &domain.User{
		Username: req.Username,
		Password: req.Password,
	}

	if user, err := s.repo.UserByName(req.Username); err != nil {
		log.Errorf("add user failed, err: %+v", err)
		c.String(http.StatusInternalServerError, "Internal Error")
		return
	} else if user != nil {
		c.String(http.StatusBadRequest, "username existed")
		return
	}

	if err := s.repo.AddUser(user); err != nil {
		log.Errorf("add user failed, err: %+v", err)
		c.String(http.StatusInternalServerError, "Internal Error")
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(CreateUserReply{UserID: user.UserID}))
}

func (s *UserService) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.BindJSON(&req); err != nil {
		log.Infof("login bad request, err: %+v", err)
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	user, err := s.repo.UserByName(req.Username)
	if err != nil {
		log.Errorf("add use failed, err: %+v", err)
		c.String(http.StatusInternalServerError, "Internal Error")
		return
	}
	if user == nil {
		c.String(http.StatusBadRequest, "user not existed")
		return
	}

	if req.Password != user.Password {
		c.JSON(http.StatusOK, FailResponse(10000, "incorrect password"))
		return
	}

	token, err := auth.GenerateToken(user.UserID, user.Username)
	if err != nil {
		log.Errorf("generate token failed, err: %+v", err)
		c.String(http.StatusInternalServerError, "Internal Error")
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(LoginReply{Token: token}))
}

func (s *UserService) FetchUser(c *gin.Context) {
	var req FetchUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	users, err := s.repo.AllUsers()
	if err != nil {
		log.Errorf("get all users failed, err: %+v", err)
		c.String(http.StatusInternalServerError, "Internal Error")
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(FetchUserReply{Users: users}))
}
