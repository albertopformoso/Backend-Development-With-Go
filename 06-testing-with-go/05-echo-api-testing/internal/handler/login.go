package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"go-api/internal/authorization"
	"go-api/internal/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {

	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		resp := newResponse(Error, "sturcutre not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if !isLoginValid(&data) {
		resp := newResponse(Error, "user or password not valid", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "can't generate token", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)

}

func isLoginValid(data *model.Login) bool {
	return data.Email == "albert@test.com" && data.Password == "123456"
}
