package exception

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateLog(exc InitialExceptionCreateResponse, code int, message string, messageInd string) {

	requestId := CreateRequestId(exc)

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		exc.Log.Info(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s"`,
			requestId,
			code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			message,
			messageInd,
		))
	} else if code == fiber.StatusBadRequest || code == fiber.StatusUnauthorized || code == fiber.StatusNotFound {
		exc.Log.Warn(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s"`,
			requestId,
			code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			message,
			messageInd,
		))
	} else {
		exc.Log.Error(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s"`,
			requestId,
			code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			message,
			messageInd,
		))
	}

}

func CreateLog_Data(exc InitialExceptionCreateResponse, code int, message string, messageInd string, data interface{}) {

	requestId := CreateRequestId(exc)

	var stringDataJson string
	if data != nil {
		dataJson, _ := json.Marshal(data)
		count := strings.Count(string(dataJson), `"`)
		countNil := strings.Count(string(dataJson), `\u003cnil\u003e`)
		stringFilterNil := strings.Replace(string(dataJson), `\u003cnil\u003e`, ``, countNil)
		stringDataJson = strings.Replace(stringFilterNil, `"`, `'`, count)
	}

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		exc.Log.Info(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s" data="%s"`,
			requestId,
			code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			message,
			messageInd,
			stringDataJson,
		))
	} else if code == fiber.StatusBadRequest || code == fiber.StatusUnauthorized || code == fiber.StatusNotFound {
		exc.Log.Warn(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s" data="%s"`,
			requestId,
			code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			message,
			messageInd,
			stringDataJson,
		))
	} else {
		exc.Log.Error(fmt.Sprintf(`requestId="%v" code="%v" method="%s" path="%v" msg="%s" msgInd="%s" data="%s"`,
			requestId,
			code,
			exc.Ctx.Method(),
			exc.Ctx.Path(),
			message,
			messageInd,
			stringDataJson,
		))
	}
}

func CreateResponse(exc InitialExceptionCreateResponse, code int, message string, messageInd string, data interface{}) error {

	requestId := CreateRequestId(exc)

	RespData := ResponseData{
		RequestId: requestId,
		Data:      data,
		Status: StatusResponseData{
			Code:       code,
			Message:    message,
			MessageInd: messageInd,
		},
		TimeStamp: time.Now(),
	}

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		RespData.Status.Type = "SUCCESS"
	} else {
		RespData.Status.Type = "ERROR"
	}

	switch code {
	case fiber.StatusOK:
		RespData.Status.Name = "OK"
	case fiber.StatusCreated:
		RespData.Status.Name = "CREATED"
	case fiber.StatusBadRequest:
		RespData.Status.Name = "BAD REQUEST"
	case fiber.StatusUnauthorized:
		RespData.Status.Name = "UNAUTHORIZED"
	case fiber.StatusNotFound:
		RespData.Status.Name = "NOT FOUND"
	case fiber.StatusInternalServerError:
		RespData.Status.Name = "INTERNAL SERVER ERROR"
	}

	return exc.Ctx.Status(code).JSON(RespData)
}

func CreateResponse_Log(exc InitialExceptionCreateResponse, code int, message string, messageInd string, data interface{}) error {

	requestId := CreateRequestId(exc)

	RespData := ResponseData{
		RequestId: requestId,
		Data:      data,
		Status: StatusResponseData{
			Code:       code,
			Message:    message,
			MessageInd: messageInd,
		},
		TimeStamp: time.Now(),
	}

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		RespData.Status.Type = "SUCCESS"
	} else {
		RespData.Status.Type = "ERROR"
	}

	switch code {
	case fiber.StatusOK:
		RespData.Status.Name = "OK"
	case fiber.StatusCreated:
		RespData.Status.Name = "CREATED"
	case fiber.StatusBadRequest:
		RespData.Status.Name = "BAD REQUEST"
	case fiber.StatusUnauthorized:
		RespData.Status.Name = "UNAUTHORIZED"
	case fiber.StatusNotFound:
		RespData.Status.Name = "NOT FOUND"
	case fiber.StatusInternalServerError:
		RespData.Status.Name = "INTERNAL SERVER ERROR"
	}

	DefaultLog(exc, requestId, code, RespData)

	return exc.Ctx.Status(code).JSON(RespData)
}

func CreateResponse_Page(exc InitialExceptionCreateResponse, code int, message string, messageInd string, data interface{}, page int, pageSize int, totalData int) error {

	requestId := CreateRequestId(exc)

	//* Get total page
	totalPage := int(math.Ceil(float64(totalData) / float64(pageSize)))
	if pageSize == 0 {
		totalPage = 1
	}

	//* Get Is next
	var isNext bool
	if totalPage-page < 1 {
		isNext = false
	} else {
		isNext = true
	}

	RespData := ResponseData{
		RequestId: requestId,
		Data:      data,
		Status: StatusResponseData{
			Code:       code,
			Message:    message,
			MessageInd: messageInd,
		},
		TimeStamp: time.Now(),
		Page: Page{
			CurrentPage: page,
			TotalPage:   totalPage,
			TotalData:   totalData,
			PageSize:    pageSize,
			IsNext:      isNext,
		},
	}

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		RespData.Status.Type = "SUCCESS"
	} else {
		RespData.Status.Type = "ERROR"
	}

	switch code {
	case fiber.StatusOK:
		RespData.Status.Name = "OK"
	case fiber.StatusCreated:
		RespData.Status.Name = "CREATED"
	case fiber.StatusBadRequest:
		RespData.Status.Name = "BAD REQUEST"
	case fiber.StatusUnauthorized:
		RespData.Status.Name = "UNAUTHORIZED"
	case fiber.StatusNotFound:
		RespData.Status.Name = "NOT FOUND"
	case fiber.StatusInternalServerError:
		RespData.Status.Name = "INTERNAL SERVER ERROR"
	}

	return exc.Ctx.Status(code).JSON(RespData)
}

func CreateResponse_Log_Page(exc InitialExceptionCreateResponse, code int, message string, messageInd string, data interface{}, page int, pageSize int, totalData int) error {

	requestId := CreateRequestId(exc)

	//* Get total page
	totalPage := int(math.Ceil(float64(totalData) / float64(pageSize)))
	if pageSize == 0 {
		totalPage = 1
	}

	//* Get Is next
	var isNext bool
	if totalPage-page < 1 {
		isNext = false
	} else {
		isNext = true
	}

	RespData := ResponseData{
		RequestId: requestId,
		Data:      data,
		Status: StatusResponseData{
			Code:       code,
			Message:    message,
			MessageInd: messageInd,
		},
		TimeStamp: time.Now(),
		Page: Page{
			CurrentPage: page,
			TotalPage:   totalPage,
			TotalData:   totalData,
			PageSize:    pageSize,
			IsNext:      isNext,
		},
	}

	if code == fiber.StatusOK || code == fiber.StatusCreated {
		RespData.Status.Type = "SUCCESS"
	} else {
		RespData.Status.Type = "ERROR"
	}

	switch code {
	case fiber.StatusOK:
		RespData.Status.Name = "OK"
	case fiber.StatusCreated:
		RespData.Status.Name = "CREATED"
	case fiber.StatusBadRequest:
		RespData.Status.Name = "BAD REQUEST"
	case fiber.StatusUnauthorized:
		RespData.Status.Name = "UNAUTHORIZED"
	case fiber.StatusNotFound:
		RespData.Status.Name = "NOT FOUND"
	case fiber.StatusInternalServerError:
		RespData.Status.Name = "INTERNAL SERVER ERROR"
	}

	DefaultLog(exc, requestId, code, RespData)

	return exc.Ctx.Status(code).JSON(RespData)
}
