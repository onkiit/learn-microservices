package users

import (
	"context"
)

type Store interface {
	GetByID(string) (*User, error)
	GetByIDContext(context.Context, string) (*User, error)
}
