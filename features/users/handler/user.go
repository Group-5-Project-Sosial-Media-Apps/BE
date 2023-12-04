package user

import (
	"net/http"
	"sosmed/features/users"
	// "sosmed/helper/jwt"
	"sosmed/helper/responses"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	srv users.Service
}

func New(s users.Service) users.Handler {
	return &userController{
		srv: s,
	}
}

func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(UserRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(users.User)
		inputProcess.Nama = input.Nama
		inputProcess.Email = input.Email
		inputProcess.Password = input.Password
		inputProcess.UserName = input.UserName

		result, err := uc.srv.Register(*inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return responses.PrintResponse(c, statusCode, message, nil)
		}

		var response = new(UserResponse)
		response.ID = result.ID
		response.Nama = result.Nama
		response.UserName = result.UserName
		response.Email = result.Email

		return responses.PrintResponse(c, http.StatusCreated, "success create data", response)
	}
}
