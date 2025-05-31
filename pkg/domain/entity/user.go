package entity

import (
	"time"
)

type UserID int64

type User struct {
	ID        UserID    `validate:"gte=1"`
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required,userUpdateAfterCreation"`
	UserUID   string    `validate:"required"`
}

func NewUser(id int64, userUID string, createdAt, updatedAt time.Time) (*User, error) {
	user := &User{
		ID:        UserID(id),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		UserUID:   userUID,
	}
	return user, nil
}
