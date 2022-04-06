package main

import (
	"backend-test/db"
	"backend-test/handler"
	"backend-test/helper"
	log "backend-test/log"
	"backend-test/repository/repo_impl"
	"backend-test/router"
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("DEV ENVIROMENT")
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

	StructValidator := helper.NewStructValidator()
	StructValidator.RegisterValidate()

	e.Validator = StructValidator

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetupRouter()
	go scheduleUpdateTrending(10000*time.Second, repoHandler)
	e.Logger.Fatal(e.Start(":3000"))
}

func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
