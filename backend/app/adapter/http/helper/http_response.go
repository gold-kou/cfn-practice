package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gold-kou/cfn-practice/backend/app/adapter/http/model"
)

func ResponseSuccess(w http.ResponseWriter) {
	resp := model.ResponseSuccess{
		Status:  http.StatusOK,
		Message: "success",
	}
	w.Header().Set(HeaderKeyContentType, HeaderValueApplicationJSON)
	// superfluousでるため、w.WriteHeader(http.StatusOK)はしない
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err.Error())
	}
}

func ResponseBadRequest(w http.ResponseWriter, message string) {
	resp := model.ResponseBadRequest{
		Status:  http.StatusBadRequest,
		Message: message,
	}
	w.Header().Set(HeaderKeyContentType, HeaderValueApplicationJSON)
	w.Header().Set(HeaderKeyCacheControl, HeaderValueNoStore)
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err.Error())
	}
}

func ResponseNotFound(w http.ResponseWriter, message string) {
	resp := model.ResponseNotFound{
		Status:  http.StatusNotFound,
		Message: message,
	}
	w.Header().Set(HeaderKeyContentType, HeaderValueApplicationJSON)
	w.Header().Set(HeaderKeyCacheControl, HeaderValueNoStore)
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err.Error())
	}
}

func ResponseNotAllowedMethod(w http.ResponseWriter, message string, methods []string) {
	resp := model.ResponseNotAllowedMethod{
		Status:  http.StatusMethodNotAllowed,
		Message: message,
	}
	w.Header().Set(HeaderKeyContentType, HeaderValueApplicationJSON)
	w.Header().Set(HeaderKeyCacheControl, HeaderValueNoStore)
	for _, m := range methods {
		w.Header().Set(HeaderKeyAllow, m)
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err.Error())
	}
}

func ResponseInternalServerError(w http.ResponseWriter, message string) {
	// MEMO Sentryのお金ないので、リクエストとの紐付けが難しいが、標準ログを出すしかない
	log.Println("internal server error: " + message)

	resp := model.ResponseInternalServerError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
	w.Header().Set(HeaderKeyContentType, HeaderValueApplicationJSON)
	w.Header().Set(HeaderKeyCacheControl, HeaderValueNoStore)
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err.Error())
	}
}
