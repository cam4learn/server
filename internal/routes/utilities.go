package routes

import (
	"net/http"
	"server/internal/registration"
	"server/internal/validator"

	"github.com/gin-gonic/gin"
)

func getInput(c *gin.Context) (string, string) {
	login, loginOK := c.GetPostForm("login")
	password, passwordOK := c.GetPostForm("password")

	if !(loginOK && passwordOK) {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	ok, _ := validator.Validate(registration.AuthData{login, password})
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return login, password
}

func adminAddDBAction(c *gin.Context, form *interface{}, action func(interface{})) {
	if c.Bind(form) == nil {
		action(*form)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}
