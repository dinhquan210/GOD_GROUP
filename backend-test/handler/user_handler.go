package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserId string
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "quan",
		"email": "quan2210quangmail.com",
	})

}

func (u *UserHandler) HandleSignUp(c echo.Context) error {
	type User struct {
		Name  string `json:"name"`  //custome ten trong struc
		Email string `json:"email"` //
		Age   int    `json:"age"`
	}

	user := User{
		Name:  "quan",
		Email: "quan123@gmail.com",
		Age:   10,
	}
	fmt.Println(user)
	err := c.JSON(http.StatusOK, user)
	return err
}
