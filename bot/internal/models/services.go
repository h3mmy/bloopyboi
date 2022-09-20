package models

type BloopyService interface {
	GetClient() interface{}
	Verify() bool
}
