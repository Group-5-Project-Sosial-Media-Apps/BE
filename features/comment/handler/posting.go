package comment

import (
	"net/http"
	"sosmed/features/comment"
	"strings"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	s comment.Service
}

func New(s comment.Service) comment.Handler {
	return &CommentHandler{
		s: s,
	}
}

func (bc *CommentHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		// userID, _ := jwt.ExtractToken(c.Get("user").(*gojwt.Token))
		var input = new(CommentRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		var inputProcess = new(comment.Comment)
		inputProcess.Pesan = input.Pesan
		inputProcess.PostID = input.PostingID

		result, err := bc.s.TambahComment(c.Get("user").(*gojwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return c.JSON(statusCode, map[string]any{
				"message": message,
			})
		}

		var responsePost = new(CommentResponse)
		responsePost.CommentID = result.ID
		responsePost.Pesan = result.Pesan
		responsePost.User.UserID = result.Users.ID
		responsePost.User.Foto = result.Users.Foto
		responsePost.User.Nama = result.Users.Nama
		responsePost.User.UserName = result.Users.UserName

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    responsePost,
		})
	}
}
