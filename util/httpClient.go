package util

type BloopyHttp struct {
	inspiro_api *InspiroClient
}

func NewBloopyHttpClient(inspiro *InspiroClient) *BloopyHttp {
	return &BloopyHttp{inspiro_api: inspiro}
}
