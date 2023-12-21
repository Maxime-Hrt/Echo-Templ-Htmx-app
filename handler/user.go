package handler

import (
	"github.com/Maxime-Hrt/got/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandlerUserShow(c echo.Context) error {
	return render(c, user.Show())
}
