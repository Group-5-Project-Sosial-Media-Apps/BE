package posting

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sosmed/features/posting"
	"strconv"
	"strings"

	cld "sosmed/utils/cloudinary"

	"github.com/cloudinary/cloudinary-go/v2"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type PostingHandler struct {
	s      posting.Service
	cl     *cloudinary.Cloudinary
	ct     context.Context
	folder string
}

func New(s posting.Service, cld *cloudinary.Cloudinary, ctx context.Context, uploadparam string) posting.Handler {
	return &PostingHandler{
		s:      s,
		cl:     cld,
		ct:     ctx,
		folder: uploadparam,
	}
}

func (bc *PostingHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(PostingRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		formHeader, _ := c.FormFile("foto")

		var link string

		if formHeader != nil {

			formFile, err := formHeader.Open()
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError, map[string]any{
						"message": "formfile error",
					})
			}

			link, err = cld.UploadImage(bc.cl, bc.ct, formFile, bc.folder)
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

			input.Foto = link

		}

		var inputProcess = new(posting.Posting)
		inputProcess.Postingan = input.Pesan
		inputProcess.Foto = link

		result, err := bc.s.TambahPosting(c.Get("user").(*gojwt.Token), *inputProcess)

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

		var responsePost = new(PostingResponse)
		responsePost.PostingID = result.ID
		responsePost.Pesan = result.Postingan
		responsePost.Foto = result.Foto
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

func (ga *PostingHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil || page <= 0 {
			page = 1
		}

		pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
		if err != nil || pageSize <= 0 {
			pageSize = 10
		}

		dataPosting, totalCount, err := ga.s.GetAllPosting(page, pageSize)
		if err != nil {
			c.Logger().Error("ERROR GetAll, explain:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Error retrieving paginated Kupons",
			})
		}

		fmt.Println(dataPosting)

		totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

		var responses []PostingResponse
		for _, result := range dataPosting {
			response := PostingResponse{
				PostingID: result.ID,
				Pesan:     result.Postingan,
				Foto:      result.Foto,
				User: PostingResponseUser{
					UserID:   result.Users.ID,
					Nama:     result.Users.Nama,
					UserName: result.Users.UserName,
					Foto:     result.Users.Foto,
				},
			}
			responses = append(responses, response)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":    "success get all data",
			"data":       responses,
			"pagination": map[string]interface{}{"page": page, "pageSize": pageSize, "totalPages": totalPages},
		})

	}
}
