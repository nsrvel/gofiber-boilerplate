package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/exception"
	"github.com/nsrvel/go-fiber-boilerplate/pkg/utils"
	"golang.org/x/exp/slices"
)

//* Roles
func JWTIsUser() fiber.Handler {
	return CheckJWT([]int{1})
}
func JWTIsAdmin() fiber.Handler {
	return CheckJWT([]int{2})
}

const (
	qSelect = `
	`
)

func IsAuthenticated(token string) error {
	var compared string
	err := initData.DBList.SimpleTransaction.Raw(qSelect, token).Scan(&compared).Error
	if err != nil {
		return errors.New("Token invalid or the account is logged in elsewhere, err: " + err.Error())
	}
	if compared != token {
		return errors.New("Token invalid or the account is logged in elsewhere, please login again")
	}
	return nil
}

func CheckJWT(functions []int) fiber.Handler {
	return func(c *fiber.Ctx) error {

		init := exception.InitException(c, initData.Conf, initData.Log)

		authorizationHeader := c.Get("Authorization")
		var token string

		if !strings.Contains(authorizationHeader, "Bearer") {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Invalid token format", "Format token tidak valid", nil)
		}
		token = strings.Replace(authorizationHeader, "Bearer ", "", -1)

		//* Check EXP
		claim, err := utils.CheckAccessToken(initData.Conf, token)
		if err != nil {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, "Token Expired", "Token kedaluwarsa", nil)
		}

		accessID := int64(claim["access_id"].(float64))
		username := string(claim["username"].(string))
		name := string(claim["name"].(string))
		isAdmin := bool(claim["is_admin"].(bool))

		//* Check validation on db
		err = IsAuthenticated(token)
		if err != nil {
			return exception.CreateResponse_Log(init, fiber.StatusUnauthorized, err.Error(), err.Error(), nil)
		}

		if slices.Contains(functions, 2) {
			if isAdmin != true {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
					"status":  "error",
					"message": "Youre not allowed to acces this endpoint",
				})
			}
		}

		c.Locals("access_id", accessID)
		c.Locals("username", username)
		c.Locals("name", name)

		return c.Next()
	}
}
