package user

import (
	entity "github.com/avinashb98/bookstore-users-api/domain/user"
	"github.com/avinashb98/bookstore-users-api/service"
	restErrors "github.com/avinashb98/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := restErrors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := service.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user entity.User

	if err := c.ShouldBindJSON(&user); err != nil {
		jsonError := restErrors.NewBadRequestError("Invalid json body")
		c.JSON(jsonError.Status, jsonError)
		return
	}

	result, saveErr := service.CreateUser(user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "Implement Me!",
	})
}
