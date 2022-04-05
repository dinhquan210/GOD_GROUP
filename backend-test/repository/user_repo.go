package repository

import (
	"backend-test/model"
	"backend-test/model/req"

	"golang.org/x/net/context"
)

type UserRepo interface { // ben trong dinh nghia cac phuong thuc de tuong tac voi table user
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SaveUser(context context.Context, user model.User) (model.User, error)
}
