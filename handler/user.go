package handler

import (
	"github.com/dan-lay/test-go-app/model"
	"github.com/dan-lay/test-go-app/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@gmail.com",
	}
	return render(c, user.Show(u))
}