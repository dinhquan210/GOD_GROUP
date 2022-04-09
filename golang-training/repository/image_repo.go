package repository

import (
	"context"
	"golang-training/model"
)

type ImageRepo interface {
	SaveImage(context context.Context, image model.Image) (model.Image, error)
}
