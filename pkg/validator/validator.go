package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/exception"
)

//? A Custom struct validation
func ValidateDataRequest(dataReq interface{}) (string, string) {
	//* Validate data request
	validate := validator.New()
	validate.RegisterValidation("confirmPassword", DefaultConfirmPasswordValidator)
	validate.RegisterValidation("date1", DefaultFormatDateValidator)
	validate.RegisterValidation("phone1", DefaultPhoneNumberValidator)
	validate.RegisterValidation("minPassword", DefaultMinimalLengthPasswordValidator)
	err := validate.Struct(dataReq)

	var errMsg string
	var errMsgInd string

	if err != nil {
		if strings.Contains(err.Error(), "failed on the 'required' tag") {
			field := between(err.Error(), "Field validation for ", " failed on the 'required")
			errMsg = fmt.Sprintf("Data %s cannot be empty", field)
			errMsgInd = fmt.Sprintf("Data %s tidak boleh kosong", field)
		} else if strings.Contains(err.Error(), "failed on the 'email' tag") {
			errMsg = "Data Email not valid"
			errMsgInd = "Data Email tidak valid"
		} else if strings.Contains(err.Error(), "failed on the 'confirmPassword' tag") {
			errMsg = "Password not equal"
			errMsgInd = "Password tidak sama"
		} else if strings.Contains(err.Error(), "failed on the 'minPassword' tag") {
			errMsg = "Password cannot be less than 8 characters"
			errMsgInd = "Kata sandi tidak boleh kurang dari 8 karakter"
		} else {
			errMsg = err.Error()
			errMsgInd = err.Error()
		}
		return errMsg, errMsgInd
	}
	return "", ""
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func DefaultFormatDateValidator(field validator.FieldLevel) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	value := field.Field().String()
	if value == "" {
		return true
	}
	return re.MatchString(value)
}

func DefaultPhoneNumberValidator(field validator.FieldLevel) bool {
	val := regexp.MustCompile(`^(\+62|62)?[\s-]?0?8[1-9]{1}\d{1}[\s-]?\d{4}[\s-]?\d{2,5}$`)
	value := field.Field().String()
	if value == "" {
		return true
	}
	return val.MatchString(value)
}

func DefaultConfirmPasswordValidator(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		return false
	}
	newPassword := field.Field().String()
	confirmPassword := value.String()
	if newPassword == "" && confirmPassword == "" {
		return true
	}
	return newPassword == confirmPassword
}

func DefaultMinimalLengthPasswordValidator(field validator.FieldLevel) bool {
	value := field.Field().String()
	if value == "" {
		return true
	}
	if len(value) < 8 {
		return false
	}
	return true
}

func DefaultUsernameValidator(username string) *exception.Error {

	length := len(username)
	firstChar := username[0:1]
	lastChar := username[length-1 : length]

	rule1 := regexp.MustCompile(`^[a-zA-Z._]*$`)
	if !rule1.MatchString(username) {
		return exception.NewError(fiber.StatusBadRequest, "Username can contain only alphabets, periods( . ), and underscores ( _ )", "Username hanya boleh alfabet, titik( . ), and garis bawah ( _ )")
	}
	rule2 := regexp.MustCompile(`^[a-zA-Z]*$`)
	if !rule2.MatchString(firstChar) {
		return exception.NewError(fiber.StatusBadRequest, "Username cannot start with a symbol", "Username tidak boleh diawali dengan symbol")
	}
	if !rule2.MatchString(lastChar) {
		return exception.NewError(fiber.StatusBadRequest, "Username cannot end with a symbol", "Username tidak boleh diakhiri dengan symbol")
	}
	if strings.Contains(username, "._") || strings.Contains(username, "_.") || strings.Contains(username, "..") || strings.Contains(username, "__") {
		return exception.NewError(fiber.StatusBadRequest, "Username cannot contain consecutive symbols", "Username tidak boleh berisi urutan symbol")
	}
	if length < 6 {
		return exception.NewError(fiber.StatusBadRequest, "Username cannot be less than 6 characters", "Username tidak boleh kurang dari 6 karakter")
	}
	if length > 30 {
		return exception.NewError(fiber.StatusBadRequest, "Username cannot be longer than 30 characters", "Username tidak boleh lebih dari 30 karakter")
	}
	return nil
}
