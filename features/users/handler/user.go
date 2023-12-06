package user

import (
	"context"
	"net/http"
	"sosmed/features/users"
	"sosmed/helper/jwt"
	"sosmed/helper/responses"
	"strings"

	cld "sosmed/utils/cloudinary"

	"github.com/cloudinary/cloudinary-go/v2"
	golangjwt "github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type userController struct {
	srv    users.Service
	cl     *cloudinary.Cloudinary
	ct     context.Context
	folder string
}

func New(s users.Service, cld *cloudinary.Cloudinary, ctx context.Context, uploadparam string) users.Handler {
	return &userController{
		srv:    s,
		cl:     cld,
		ct:     ctx,
		folder: uploadparam,
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

func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uc.srv.Login(input.UserName, input.Password)

		if err != nil {
			c.Logger().Error("ERROR Login, explain:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "data yang diinputkan tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika memproses data",
			})
		}

		strToken, err := jwt.GenerateJWT(result.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "terjadi permasalahan ketika mengenkripsi data",
			})
		}

		var response = new(LoginResponse)
		response.Nama = result.Nama
		response.ID = result.ID
		response.Token = strToken

		return c.JSON(http.StatusOK, map[string]any{
			"message": "login success",
			"data":    response,
		})
	}
}

func (uc *userController) GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(GetUserByIdRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uc.srv.GetUserById(input.ID)
		if err != nil || input.ID != result.ID || input.ID == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "user tidak ditemukan",
			})
		}
		var response = new(GetUserByIdResponse)
		response.ID = result.ID
		response.Nama = result.Nama
		response.UserName = result.UserName
		response.Email = result.Email
		response.Foto = result.Foto

		return c.JSON(http.StatusOK, map[string]any{
			"message": "get user by userID successful",
			"data":    response,
		})
	}
}

func (uc *userController) DelUserById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(DelUserByIdRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := uc.srv.DelUserById(input.ID)
		if err != nil || input.ID != result.ID || input.ID == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "user tidak ditemukan",
			})
		}
		var response = new(DelUserByIdResponse)
		response.ID = result.ID
		response.Nama = result.Nama
		response.UserName = result.UserName
		response.Email = result.Email

		return c.JSON(http.StatusOK, map[string]any{
			"message": "delete user by userID successful",
			"data":    response,
		})
	}
}


func (up *userController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := jwt.ExtractToken(c.Get("user").(*golangjwt.Token))
		var input = new(UserUpdate)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
			})
		}

		formHeader, err := c.FormFile("foto")
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError, map[string]any{
					"messege": "formheader error",
				})
		}

		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError, map[string]any{
					"message": "formfile error",
				})
		}

		link, err := cld.UploadImage(up.cl, up.ct, formFile, up.folder)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "harap pilih gambar",
					"data":    nil,
				})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "kesalahan pada server",
					"data":    nil,
				})
			}
		}

		// var update = link

		input.Foto = link

		updatedUser := users.User{
			Nama:     input.Nama,
			UserName: input.UserName,
			Foto:     input.Foto,
		}

		result, err := up.srv.UpdateUser(userID, updatedUser)
		if err != nil {
			c.Logger().Error("ERROR UpdateUser, explain:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to update user",
			})
		}

		result.Foto = link

		var response = &UserUpdate{
			ID:       result.ID,
			Nama:     result.Nama,
			UserName: result.UserName,
			Foto:     result.Foto,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "user updated successfully",
			"data":    response,
		})

	}

}
