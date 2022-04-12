package repo_impl

import (
	"context"
	"fmt"
	"golang-training/banana"
	"golang-training/db"
	"golang-training/model"
	"golang-training/repository"
	"time"
)

type ImageRepoImpl struct {
	sql *db.Sql
}

func NewImageRepo(sql *db.Sql) repository.ImageRepo {
	return &ImageRepoImpl{
		sql: sql,
	}
}

func (i *ImageRepoImpl) SaveImage(context context.Context, image model.Image) (model.Image, error) {
	statement := `
	INSERT INTO images(id, urls_full, urls_raw, urls_regular, created_at, updated_at, width,height)
	VALUES(:id, :urls_full, :urls_raw, :urls_regular, :created_at, :updated_at, :width, :height)
`
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now()
	_, err := i.sql.Db.NamedExecContext(context, statement, image)
	if err != nil {
		fmt.Println(err)
		return image, banana.GetRandomFail
	}
	return image, err
}
