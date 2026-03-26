package controllers

import (
	"edart-core/app/dtos"
	"edart-core/app/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service * services.UserService
}

func NewUserController() * UserController {
	return &UserController{service: services.NewUserService()}
}

func (c *UserController) CreateNewUser(ctx *gin.Context) {
	var request dtos.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid Request Body",
		})
		return
	}

	user, err := c.service.CreateUserService(&request)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Failed to create user",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
			"uid":       user.Uid,
			"fullname":  user.Fullname,
			"country":   user.Country,
			"dob":       user.Dob.Format("2006-01-02"),
			"createdAt": user.CreateAt.Format("2006-01-02 15:04:05"),
		},)
}