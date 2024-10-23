package ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/tienphuckx/ecom-backbone-api.git/internal/service"
	"github.com/tienphuckx/ecom-backbone-api.git/pkg/response"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("id")
	userInfo := uc.UserService.GetUser(userID)
	response.ServerResponseSuccess(c, response.ErrCode_SUCCESS, userInfo)
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	userEmail := c.Param("email")
	userInfo := uc.UserService.GetUserByEmail(userEmail)
	response.ServerResponseSuccess(c, response.ErrCode_SUCCESS, userInfo)
}
