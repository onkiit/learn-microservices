package auth

import (
	"log"

	"github.com/kataras/iris/v12"
)

func (s *Services) Login(ctx iris.Context) {
	user, err := s.userStore.GetByID("123")
	if err != nil {
		return
	}

	log.Println("found user", user)
}
