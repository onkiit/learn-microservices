package auth

import (
	"github.com/kataras/iris/v12"
	"github.com/onkiit/oserver/api"
)

type Services struct{}

func (s *Services) RegisterRouter(router iris.Party) {
	router.Get("/login", s.Login)
}

func New() api.Router {
	return &Services{}
}

func init() {
	api.RegisterRouters(New())
}
