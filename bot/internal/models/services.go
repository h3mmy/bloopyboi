package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type BloopyService interface {
	GetClient() interface{}
	Verify() bool
}

type BloopyMeta struct {
	Id        uuid.UUID
	CreatedAt time.Time
	OwnerRef  []string
}

func NewBloopyMeta(owners ...string) BloopyMeta {
	return BloopyMeta{
		Id: uuid.New(),
		CreatedAt: time.Now(),
		OwnerRef: owners,
	}
}
