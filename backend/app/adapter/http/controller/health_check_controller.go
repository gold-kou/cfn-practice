package controller

import (
	"fmt"
	"net/http"

	"github.com/gold-kou/cfn-practice/backend/app/adapter/http/helper"
	"github.com/gold-kou/cfn-practice/backend/app/adapter/mysql"
)

func HealthController(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/health/liveness":
		switch r.Method {
		case http.MethodGet:
			helper.ResponseSuccess(w)
		default:
			methods := []string{http.MethodGet}
			helper.ResponseNotAllowedMethod(w, errMsgNotAllowedMethod, methods)
		}
	case "/health/readiness":
		switch r.Method {
		case http.MethodGet:
			err := getHealthReadiness()
			if err != nil {
				helper.ResponseInternalServerError(w, fmt.Sprintf("readiness error: %s", err.Error()))
			}
			helper.ResponseSuccess(w)
		default:
			methods := []string{http.MethodGet}
			helper.ResponseNotAllowedMethod(w, errMsgNotAllowedMethod, methods)
		}
	default:
		helper.ResponseInternalServerError(w, errMsgControllerPath)
	}
}

func getHealthReadiness() error {
	db, err := mysql.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
