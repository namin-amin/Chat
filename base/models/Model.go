package models

import (
	"github.com/labstack/echo/v4"
)

// IBaseController
//
// interface defines the basic Methods to be implemented by a controller
type IBaseController interface {
	RegisterController(e *echo.Group)
}
