package exception

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/go-fiber-boilerplate/config"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ResponseData struct {
	RequestId string             `json:"requestId"`
	Data      interface{}        `json:"data"`
	Status    StatusResponseData `json:"status"`
	TimeStamp time.Time          `json:"timeStamp"`
	Page      interface{}        `json:"page,omitempty"`
}

type Page struct {
	CurrentPage int  `json:"currentPage"`
	TotalPage   int  `json:"totalPage"`
	TotalData   int  `json:"totalData"`
	PageSize    int  `json:"pageSize"`
	IsNext      bool `json:"isNext"`
}

type StatusResponseData struct {
	Code       int    `json:"code"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	MessageInd string `json:"messageInd"`
}

type ResponseMessageData struct {
	Message string `json:"message"`
}

type Response interface{}

type Error struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	MessageInd string `json:"messageInd"`
}

type InitialExceptionCreateResponse struct {
	Ctx  *fiber.Ctx
	Conf *config.Config
	Log  *logrus.Logger
}

func NewError(code int, message string, messageInd string) *Error {
	return &Error{
		Code:       code,
		Message:    message,
		MessageInd: messageInd,
	}
}

func InitException(c *fiber.Ctx, conf *config.Config, log *logrus.Logger) InitialExceptionCreateResponse {
	init := InitialExceptionCreateResponse{
		Ctx:  c,
		Log:  log,
		Conf: conf,
	}
	return init
}

func CreateRequestId(exc InitialExceptionCreateResponse) string {

	requestId := fmt.Sprintf("%v", exc.Ctx.Locals("requestId"))
	if requestId == "<nil>" {

		//* Generate New Request Id
		newUUID, _ := utils.GenerateUUID()
		username := fmt.Sprintf("%v", exc.Ctx.Locals("username"))
		if username == "<nil>" {
			username = strings.ToLower(exc.Conf.App.Name)
		}
		formattedUsername := strings.Replace(username, "@kalbenutritionals.com", "", 1)
		newRequestId := fmt.Sprintf("%s-%s", formattedUsername, newUUID)

		exc.Ctx.Locals("requestId", newRequestId)
		return newRequestId
	} else {
		return requestId
	}
}

func DefaultLog(exc InitialExceptionCreateResponse, requestId string, code int, respData ResponseData) {

	//* Replace " to ' in msg & msgInd if exist
	var msg string
	var msgInd string
	msg = strings.ReplaceAll(respData.Status.Message, `"`, `'`)
	msgInd = strings.ReplaceAll(respData.Status.MessageInd, `"`, `'`)

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		exc.Log.Info(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s"`,
			requestId,
			respData.Status.Code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			msg,
			msgInd,
		))
	} else if code == fiber.StatusBadRequest || code == fiber.StatusUnauthorized || code == fiber.StatusNotFound {
		exc.Log.Warn(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s"`,
			requestId,
			respData.Status.Code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			msg,
			msgInd,
		))
	} else {
		exc.Log.Error(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s"`,
			requestId,
			respData.Status.Code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			msg,
			msgInd,
		))
	}
}
