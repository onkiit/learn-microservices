package auth

import (
	"github.com/kataras/iris/v12"
	"github.com/onkiit/learn-microservices/user-service/models/users"
	"github.com/onkiit/sapp/api"
	"github.com/onkiit/sapp/config"
	"github.com/onkiit/sapp/reg"
)

type Services struct {
	cfg       *config.Configuration
	userStore users.Store
}

func (s *Services) RegisterRouter(router iris.Party) {
	router.Get("/login", s.Login)
}

func New(cfg *config.Configuration) api.Router {
	userStore := users.NewUserDummy("dummy")
	return &Services{
		cfg:       cfg,
		userStore: userStore,
	}
}

func init() {
	reg.RegisterRouters(New)
}
