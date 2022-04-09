package handler

import (
	"fmt"
	"golang-training/model"
	"golang-training/repository/repo_impl"
	"golang-training/utils/unsplashutils"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	ImageRepo repo_impl.ImageRepoImpl
}

const ()

func (i *ImageHandler) RandomImage(e echo.Context) error {
	// Create a Resty Client
	client := resty.New()
	reBody := unsplashutils.ResultType{}
	client.R().SetResult(&reBody).
		Get("https://api.unsplash.com/photos/random/?client_id=05qCv0koWY-_KqKyyCRmtrBqtbBISysGPznnA6wCNNg")

	fmt.Println(reBody)
	return e.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       reBody,
	})
}
