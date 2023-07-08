package fileUpload

import (
	baseModel "Chat/base/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IAttachmentController interface {
	IAttachmentService
	baseModel.IBaseController
}

type AttachmentController struct {
	IAttachmentService
}

func BaseControllers(e *echo.Group, controller AttachmentController) {
	//e.Use(auth.JWTAuth)

	e.POST("/single", controller.SaveSingleFile)
	e.POST("/multiple", controller.SaveManyFiles)

}

func (fc *AttachmentController) SaveSingleFile(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.NoContent(404)
	}

	id := c.Request().Header.Get("userId")

	_, err = fc.SaveFile(file, id, c.Request().Context())
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Could not save the sent file",
		}
	}
	return c.String(200, "ow")
}

func (fc *AttachmentController) SaveManyFiles(c echo.Context) error {
	return nil
}

func (fc *AttachmentController) RegisterController(e *echo.Group) {
	BaseControllers(e, *fc)
}

func NewFileController(fileService IAttachmentService) *AttachmentController {
	return &AttachmentController{
		IAttachmentService: fileService,
	}
}
