package chats

import (
	"Chat/auth"
	"Chat/chats/models"
	"Chat/sse"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IChatController interface {
	RegisterController(e *echo.Group, h *sse.Hub)
	IChatService
}

type ChatController struct {
	IChatService
}

func (cc *ChatController) BaseControllers(chatRoute *echo.Group, controller *ChatController, hub *sse.Hub) {

	chatRoute.GET("/sse", sse.RegClients(hub)) //Todo implement rate limiter

	chatRoute.POST("/sse/verify", func(c echo.Context) error {
		userinfo := new(sse.ClientVerifyDto)
		if err := c.Bind(&userinfo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request") //TODO add new binding error
		}
		hub.VerifyClient <- userinfo
		fmt.Println("verified user")
		return c.NoContent(http.StatusOK)

	}, auth.JWTAuth)

	chatRoute.POST("/message", func(c echo.Context) error {
		fmt.Println("init app" + strconv.Itoa(len(hub.Clients)))
		m := new(models.NewChatDto)
		if err := c.Bind(&m); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "bad request") //TODO add new binding error
		}
		newChat, err := controller.NewChat(c.Request().Context(), &models.Chat{
			Message:        m.Message,
			ConversationId: m.ConversationId,
		})

		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "chat sent not confirmed",
			}
		}
		chat, _ := json.Marshal(newChat)
		m.Data = string(chat)
		hub.SendMsg <- &m.Message
		fmt.Println(m)

		return c.JSON(http.StatusAccepted, newChat)
	}, auth.JWTAuth)

	chatRoute.GET("/paginate/", cc.getConversationsWithPagination, auth.JWTAuth)
}

func (cc *ChatController) getConversationsWithPagination(c echo.Context) error {
	pageCount, err1 := strconv.ParseInt(c.QueryParam("pagecount"), 0, 32)
	itemsPerPage, err2 := strconv.ParseInt(c.QueryParam("itemsperpage"), 0, 32)
	convId := c.QueryParam("convid")
	offset := -1
	if err1 != nil || err2 != nil {
		fmt.Println(pageCount)
		fmt.Println(itemsPerPage)
		fmt.Println(convId)
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "parameters sent are not valid",
		}
	}

	if pageCount > 1 {
		offset = (int(pageCount) - 1) * int(itemsPerPage)
	}
	var chats []models.Chat

	result := cc.GetDb().
		WithContext(c.Request().Context()).
		Offset(offset).
		Where("conversationId = ?", convId).Find(&chats)

	if result.Error != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "could not get the results",
		}
	}

	return c.JSON(http.StatusOK, chats)

}

// RegisterController
//
// Register all the routes for the controller and register with the app
func (cc *ChatController) RegisterController(e *echo.Group, h *sse.Hub) {
	cc.BaseControllers(e, cc, h)
}

// NewChatController
//
// Returns the instance of chatController
func NewChatController(chatService IChatService) *ChatController {
	fmt.Println("created new chatController")
	return &ChatController{
		IChatService: chatService,
	}
}
