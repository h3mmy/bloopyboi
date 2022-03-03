package util

type BloopyHttp struct {
	Inspiro_api *InspiroClient
}

func NewBloopyHttpClient(inspiro *InspiroClient) *BloopyHttp {
	return &BloopyHttp{Inspiro_api: inspiro}
}

func NewBloopyClient() *BloopyHttp {
	return NewBloopyHttpClient(NewInspiroClient())
}
