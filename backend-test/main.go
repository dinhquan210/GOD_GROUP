package main

import (
	"backend-test/db"
	log "backend-test/log"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
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
	e.Logger.Fatal(e.Start(":3000"))
}
