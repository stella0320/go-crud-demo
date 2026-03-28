package handler

import (
	"strconv"

	"go-crud-demo/model"

	"go-crud-demo/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

// handler
func (h *UserHandler) GetUsers(c *gin.Context) {
	allUsers := h.userService.GetAllUsers()
	c.JSON(200, allUsers)
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newUser := h.userService.CreateNewUser(req.Name)
	c.JSON(201, newUser)
}

func (h *UserHandler) QueryUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	result, err := h.userService.GetUserById(id)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, result)

}
