package conversations

import (
	"Chat/auth"
	baseModel "Chat/base/models"
	"Chat/conversations/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IConvController interface {
	IConvService
	baseModel.IBaseController
}

type ConvController struct {
	IConvService
}

/*
	BaseControllers creates the main routes for the Conversation API.

e: an echo Group router instance.

controller: a ConvController struct that handles Conversation routes.
*/
func BaseControllers(e *echo.Group, controller ConvController) {
	e.Use(auth.JWTAuth)
	e.POST("/", controller.createConversation)
	e.GET("/sender/:senderId", controller.getConversations)
	e.GET("/:id", controller.getConversationById)
}

func (cc *ConvController) createConversation(c echo.Context) error {
	var conv models.ConvCreateDto
	if err := c.Bind(&conv); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad request") //TODO add new binding error
	}

	newConversation, err := cc.CreateNewConv(conv.SenderId, conv.ReceiverId)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "conversation could not be created try again",
		}
	}

	return c.JSON(http.StatusAccepted, newConversation)
}

func (cc *ConvController) getConversations(c echo.Context) error {
	senderReceiverId := c.Param("senderId")
	fmt.Println(senderReceiverId)
	if senderReceiverId == "" {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "no sender id sent",
		}
	}

	conversation, err := cc.GetConv(senderReceiverId)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "could not find the conv",
		}
	}
	return c.JSON(http.StatusOK, conversation)
}

func (cc *ConvController) getConversationById(c echo.Context) error {
	id := c.Param("id")

	returnVal, err := cc.GetOne(c.Request().Context(), id)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "could not get conversations",
		}
	}

	return c.JSON(http.StatusOK, returnVal)
}

func (cc *ConvController) RegisterController(e *echo.Group) {
	BaseControllers(e, *cc)
}

func NewConvController(convService IConvService) *ConvController {
	return &ConvController{
		IConvService: convService,
	}
}
