package model

type ResponseBadRequest struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}
