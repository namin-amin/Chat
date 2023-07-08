package users

import (
	"Chat/auth"
	baseModel "Chat/base/models"
	"Chat/users/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	IUserService
	baseModel.IBaseController
}

// UserController
//
// Controller implementation for Users
type UserController struct {
	IUserService
}

// BaseControllers
//
// Initialize all user routes
func BaseControllers(userRoute *echo.Group, controller UserController) {

	//SignUp new user
	userRoute.POST("/signup", controller.signUpHandle)
	userRoute.POST("/signin", controller.signInHandle)

	userRoute.Use(auth.JWTAuth)

	userRoute.POST("/isValid", func(c echo.Context) error {
		return c.NoContent(http.StatusAccepted)
	})

	userRoute.GET("/", func(c echo.Context) error {
		users, err := controller.GetAll(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, users)

	})

}

// Todo can make a generic for signIn and signOut
func (uc *UserController) signUpHandle(c echo.Context) error {
	fmt.Println("signUp")
	var signUpDto models.SignUpDto
	err := c.Bind(&signUpDto)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "supplied data invalid",
		}
	}
	returnData, err := uc.Create(signUpDto, c.Request().Context())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusAccepted, returnData)
}

// TODO fix the issue
func (uc *UserController) signInHandle(c echo.Context) error {
	fmt.Println("sign in attempted")
	var signInDto models.SigInDto
	err := c.Bind(&signInDto)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "supplied data invalid",
		}
	}

	returnData, err := uc.SignIn(signInDto, c.Request().Context())

	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusAccepted, returnData)
}

// RegisterController
//
// Register all the routes for the controller and register with the app
func (uc *UserController) RegisterController(e *echo.Group) {
	BaseControllers(e, *uc)
}

// NewUserController
//
// Returns the instance of UserController
func NewUserController(userService IUserService) *UserController {
	return &UserController{
		IUserService: userService,
	}
}
