package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lgomez-dev/GoProject-Template/view/user"
)

type UserHandler struct{}

func (h *UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
