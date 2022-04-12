package handler

import (
	"golang-training/log"
	"golang-training/model"
	"golang-training/repository"
	"golang-training/utils/unsplashutils"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	ImageRepo repository.ImageRepo
}

func (i *ImageHandler) RandomImage(c echo.Context) error {
	// Create a Resty Client
	client := resty.New()
	reBody := unsplashutils.ResultType{}
	client.R().SetResult(&reBody).
		Get("https://api.unsplash.com/photos/random/?client_id=05qCv0koWY-_KqKyyCRmtrBqtbBISysGPznnA6wCNNg")

	image := model.Image{
		ImageID:      reBody.ImageID,
		URLs_full:    reBody.URLs.URLs_full,
		URLs_regular: reBody.URLs.RULs_regular,
		URLs_Raw:     reBody.URLs.URLs_Raw,
		Width:        reBody.Width,
		Height:       reBody.Height,
		Description:  reBody.Description,
	}

	image, err := i.ImageRepo.SaveImage(c.Request().Context(), image)
	if err != nil {
		log.Error(err.Error())
		// return c.JSON(http., model.Response{
		// 	StatusCode: http.StatusConflict,
		// 	Message:    err.Error(),
		// 	Data:       nil,
		// })
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "lấy ảnh thành công",
		Data:       image,
	})
}

