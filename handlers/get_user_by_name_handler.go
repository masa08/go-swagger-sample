package handlers

import (
	"go-swagger-example/restapi/operations/user"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

type userByNameImpl struct{}

func NewGetUserByNameHandler() user.GetUserByNameHandler {
	return &userByNameImpl{}
}

func (impl *userByNameImpl) Handle(p user.GetUserByNameParams) middleware.Responder {
	log.Println("hello")
	return user.NewGetUserByNameOK()
}
