package controllers

import (
	"Chat/chats"
	"Chat/conversations"
	"Chat/fileUpload"
	"Chat/sse"

	"Chat/users"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Base struct {
	userService          users.IUserService
	userRepo             users.IUserRepo
	userController       users.IUserController
	chatService          chats.IChatService
	chatRepo             chats.IChatRepo
	chatController       chats.IChatController
	convService          conversations.IConvService
	convRepo             conversations.IConvRepo
	convController       conversations.IConvController
	attachmentService    fileUpload.IAttachmentService
	attachmentRepo       fileUpload.IAttachmentRepo
	attachmentController fileUpload.IAttachmentController
}

// InitControllers
//
// instantiate controllers here in the below format
//
//	Controller(Service(Repository))
//	RegisterController call these methods with proper group
//	as all Services and Repo are declared here can be passed around if required
func (b *Base) InitControllers(db *gorm.DB, e *echo.Echo, h *sse.Hub) {
	//Initialize user Controllers
	b.userRepo = users.NewUserRepo(db)
	b.userService = users.NewUserService(b.userRepo)
	b.userController = users.NewUserController(b.userService)
	b.userController.RegisterController(e.Group("/user"))

	//Initialize chat Controllers
	b.chatRepo = chats.NewChatRepo(db)
	b.chatService = chats.NewChatService(b.chatRepo)
	b.chatController = chats.NewChatController(b.chatService)
	b.chatController.RegisterController(e.Group("/chats"), h) //routes are protected

	//Initialize conversation Controllers
	b.convRepo = conversations.NewConvRepo(db)
	b.convService = conversations.NewConvService(b.convRepo)
	b.convController = conversations.NewConvController(b.convService)
	b.convController.RegisterController(e.Group("/conversation"))

	//Initialize File Controller
	b.attachmentRepo = fileUpload.NewFileRepo(db)
	b.attachmentService = fileUpload.NewFileService(b.attachmentRepo)
	b.attachmentController = fileUpload.NewFileController(b.attachmentService)
	b.attachmentController.RegisterController(e.Group("/fileUpload"))

	//Initialize other controller
}

func NewBase() *Base {
	return &Base{}
}
