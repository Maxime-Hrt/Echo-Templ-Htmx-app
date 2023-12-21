package handler

import (
	"github.com/Maxime-Hrt/got/model"
	"github.com/Maxime-Hrt/got/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandlerUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@example.com",
	}
	return render(c, user.Show(u))
}
