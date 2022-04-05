package main

import (
	"backend-test/db"
	"backend-test/handler"
	log "backend-test/log"
	"backend-test/repository/repo_impl"
	"backend-test/router"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     5433,
		UserName: "postgres",
		PassWord: "123456",
		DbName:   "golang",
	}

	sql.Connect()
	defer sql.Close()
	e := echo.New()
	UserHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: UserHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
