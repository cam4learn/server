package routes

import (
	"errors"
	"net/http"
	"server/internal/registration"
	"server/internal/validator"

	"github.com/gin-gonic/gin"
)

type login struct {
	Login    string `json: "login"`
	Password string `json:"password" `
}

func getInput(c *gin.Context) (string, string, error) {
	//login, loginOK := c.GetPostForm("login")
	var data login
	if c.Bind(&data) != nil {
		return "", "", errors.New("MISSING_FIELDS")
	}
	//password, passwordOK := c.GetPostForm("password")

	//if !(loginOK && passwordOK) {
	//	return "", "", errors.New("MISSING_FIELDS")
	//}
	ok, _ := validator.Validate(registration.AuthData{data.Login, data.Password})
	if !ok {
		return "", "", errors.New("INCORRECT_DATA")
	}
	return data.Login, data.Password, nil
}

func adminAddDBAction(c *gin.Context, form *interface{}, action func(interface{})) {
	if c.Bind(form) == nil {
		action(*form)
		c.Status(http.StatusOK)
		return
	}
	c.AbortWithStatus(http.StatusBadRequest)
}
