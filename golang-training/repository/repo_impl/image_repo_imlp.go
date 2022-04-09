package repo_impl

import (
	"context"
	"golang-training/db"
	"golang-training/model"
)

type ImageRepoImpl struct {
	sql *db.Sql
}

// func NewImageRepo(sql *db.Sql) repository.ImageRepo {
// 	return &ImageRepoImpl{
// 		sql: sql,
// 	}
// }

func (i *ImageRepoImpl) SaverImage(context context.Context, image model.Image) error {
	return nil
}
