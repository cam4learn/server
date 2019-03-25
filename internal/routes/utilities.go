package routes

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

func getInput(c *gin.Context) (string, string) {
	login, loginOK := c.GetPostForm("login")
	password, passwordOK := c.GetPostForm("password")

	if !(loginOK && passwordOK) {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if !validateInput(login, password) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return login, password
}

func validateInput(login, password string) bool {
	if len(login) < 4 || len(password) < 6 {
		return false
	}
	upperCaseLetter := false
	numericChars := false
	for _, c := range password {
		if unicode.IsUpper(c) {
			upperCaseLetter = true
		}
		if unicode.IsDigit(c) {
			numericChars = true
		}
	}
	return upperCaseLetter && numericChars
}

func adminAddDBAction(c *gin.Context, form *interface{}, action func(interface{})) {
	if c.Bind(form) == nil {
		action(*form)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}
