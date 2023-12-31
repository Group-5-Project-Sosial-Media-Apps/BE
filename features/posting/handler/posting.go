package posting

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sosmed/features/posting"
	"sosmed/helper/jwt"
	"strconv"
	"strings"

	cld "sosmed/utils/cloudinary"

	"github.com/cloudinary/cloudinary-go/v2"
	golangjwt "github.com/golang-jwt/jwt/v5"
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

		result, err := bc.s.TambahPosting(c.Get("user").(*golangjwt.Token), *inputProcess)

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

			for _, v := range result.Comment {
				response.Comment = append(response.Comment, CommentResponse{
					CommentID: v.ID,
					Pesan:     v.Pesan,
					User: PostingResponseUser{
						UserID:   v.Users.ID,
						Nama:     v.Users.Nama,
						UserName: v.Users.UserName,
						Foto:     v.Users.Foto,
					},
				})

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

func (gp *PostingHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, _ := jwt.ExtractToken(c.Get("user").(*golangjwt.Token))

		results, err := gp.s.GetPostingById(uint(userID))
		if err != nil {
			c.Logger().Error("ERROR GetByID, explain:", err.Error())

			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "Posting not found",
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Error retrieving Posting by ID",
			})
		}

		var response []PostingResponse
		for _, v := range results {
			response = append(response, PostingResponse{
				PostingID: v.ID,
				Pesan:     v.Postingan,
				User: PostingResponseUser{
					UserID:   v.UserID,
					Nama:     v.Users.Nama,
					UserName: v.Users.UserName,
				},
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data by ID",
			"data":    response,
		})
	}
}

func (up *PostingHandler) UpdatePosting() echo.HandlerFunc {
	return func(c echo.Context) error {

		var input = new(PostingUpdate)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid input",
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

			link, err = cld.UploadImage(up.cl, up.ct, formFile, up.folder)
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

		updatedPosting := posting.Posting{
			Postingan: input.Posting,
			Foto:      input.Foto,
		}
		result, err := up.s.UpdatePosting(input.PostingID, updatedPosting)
		if err != nil {
			c.Logger().Error("ERROR UpdatePosting, explain:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "failed to update posting",
			})
		}

		result.Foto = link

		var response = &PostingUpdate{
			PostingID: input.PostingID,
			Posting:   result.Postingan,
			Foto:      result.Foto,
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "posting updated successfully",
			"data":    response,
		})
	}
}

func (dp *PostingHandler) DelPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(DelPost)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		result, err := dp.s.DelPost(input.PostID)
		if err != nil || result.ID != input.PostID {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "post tidak ditemukan ygy",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "delete user by userID successful",
		})
	}
}
