package users

import (
	"context"
)

type dummy struct {
	conn string
}

func (d *dummy) GetByID(userID string) (*User, error) {
	return d.GetByIDContext(context.Background(), userID)
}

func (d *dummy) GetByIDContext(ctx context.Context, userID string) (*User, error) {
	user := &User{
		UserID:   "123",
		Username: "tikno",
		Name:     "Tikno",
		Email:    "tikno.edu@gmail.com",
	}

	return user, nil
}

func NewUserDummy(conn string) Store {
	return &dummy{conn: conn}
}
