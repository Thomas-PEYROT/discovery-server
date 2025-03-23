package dto

type RegisterMicroserviceBody struct {
	Name string `json:"name"`
}

type RegisterMicroserviceResponse struct {
	UUID string `json:"uuid"`
	Port uint32 `json:"port"`
}
