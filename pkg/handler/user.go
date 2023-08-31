package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRequestBody struct {
	Name string
}

// @Summary Create new user
// @Tags users
// @Description create new user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param requestBody body UserRequestBody true "user's name"
// @Success 200 {integer} json "message: User created"
// @Failure 400 {integer} json "message: Failed"
// @Router /api/users [post]
func (h *Handler) createUser(c *gin.Context) {

	var requestBody UserRequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		panic(err)
	}

	err := h.services.User.Create(requestBody.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
		})
	}

}
