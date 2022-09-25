package services

type InspiroClient struct {
	Inspiro_api *InspiroService
}

func NewInspiroHttpClient(inspiro *InspiroService) *InspiroClient {
	return &InspiroClient{Inspiro_api: inspiro}
}

func NewInspiroClient() *InspiroClient {
	return NewInspiroHttpClient(NewInspiroService())
}
