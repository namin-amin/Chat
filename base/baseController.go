package base

import (
	"github.com/labstack/echo/v4"
)

type IBaseController interface {
	RegisterController(e *echo.Group)
}
