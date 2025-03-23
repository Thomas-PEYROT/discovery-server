package dto

type UnregisterMicroserviceBody struct {
	UUID string `json:"uuid"`
}

type UnregisterMicroserviceResponse struct {
	Message string `json:"message"`
}
