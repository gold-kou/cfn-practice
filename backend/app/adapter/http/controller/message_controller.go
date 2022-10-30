package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gold-kou/cfn-practice/backend/app/adapter/http/helper"
	modelHTTP "github.com/gold-kou/cfn-practice/backend/app/adapter/http/model"
	"github.com/gold-kou/cfn-practice/backend/app/adapter/impl/repository"
	"github.com/gold-kou/cfn-practice/backend/app/adapter/mysql"
	"github.com/gold-kou/cfn-practice/backend/app/application/usecase"
	"github.com/gold-kou/cfn-practice/backend/app/domain/model"
	"github.com/gorilla/mux"
)

func MessageController(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/messages":
		switch r.Method {
		case http.MethodPost:
			err := controlRegisterMessage(r)
			switch err := err.(type) {
			case nil:
				helper.ResponseSuccess(w)
			case *helper.BadRequestError:
				helper.ResponseBadRequest(w, err.Error())
			case *helper.InternalServerError:
				helper.ResponseInternalServerError(w, err.Error())
			default:
				helper.ResponseInternalServerError(w, err.Error())
			}
		default:
			methods := []string{http.MethodPost}
			helper.ResponseNotAllowedMethod(w, errMsgNotAllowedMethod, methods)
		}
	case strings.HasPrefix(r.URL.Path, "/messages/"):
		switch r.Method {
		case http.MethodGet:
			message, err := controlGetMessage(r)
			switch err := err.(type) {
			case nil:
				resp := modelHTTP.ResponseGetMessage{
					ID:      message.ID,
					Message: message.Message,
				}
				w.Header().Set(helper.HeaderKeyContentType, helper.HeaderValueApplicationJSON)
				w.WriteHeader(http.StatusOK)
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					log.Println(err.Error())
				}
			case *helper.BadRequestError:
				helper.ResponseBadRequest(w, err.Error())
			case *helper.NotFoundError:
				helper.ResponseNotFound(w, err.Error())
			case *helper.InternalServerError:
				helper.ResponseInternalServerError(w, err.Error())
			default:
				helper.ResponseInternalServerError(w, err.Error())
			}
		case http.MethodDelete:
			err := controlDeleteMessage(r)
			switch err := err.(type) {
			case nil:
				helper.ResponseSuccess(w)
			case *helper.BadRequestError:
				helper.ResponseBadRequest(w, err.Error())
			case *helper.InternalServerError:
				helper.ResponseInternalServerError(w, err.Error())
			default:
				helper.ResponseInternalServerError(w, err.Error())
			}
		default:
			methods := []string{http.MethodGet, http.MethodDelete}
			helper.ResponseNotAllowedMethod(w, errMsgNotAllowedMethod, methods)
		}
	default:
		helper.ResponseInternalServerError(w, errMsgControllerPath)
	}
}

func controlRegisterMessage(r *http.Request) error {
	// get request parameter
	var reqRegisterMessage modelHTTP.RequestRegisterMessage
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return helper.NewBadRequestError(err.Error())
	}
	defer r.Body.Close()
	if err := json.Unmarshal(b, &reqRegisterMessage); err != nil {
		return helper.NewBadRequestError(err.Error())
	}

	// validation check
	err = reqRegisterMessage.ValidateParam()
	if err != nil {
		return helper.NewBadRequestError(err.Error())
	}

	// new db
	db, err := mysql.NewDB()
	if err != nil {
		return helper.NewInternalServerError(err.Error())
	}
	defer db.Close()

	// UseCase
	u := usecase.NewRegisterMessage(reqRegisterMessage, repository.NewMessageRepository(db))
	if err = u.RegisterMessageUseCase(r.Context()); err != nil {
		// duplicate errorも500とする
		return helper.NewInternalServerError(err.Error())
	}
	return err
}

func controlGetMessage(r *http.Request) (message model.Message, err error) {
	// get parameter
	vars := mux.Vars(r)
	paramMessageID, ok := vars["id"]
	if !ok {
		return message, helper.NewBadRequestError("id is required")
	}
	messageID, err := strconv.Atoi(paramMessageID)
	if err != nil {
		return message, helper.NewInternalServerError(err.Error())
	}

	// db new
	db, err := mysql.NewDB()
	if err != nil {
		err = helper.NewInternalServerError(err.Error())
		return
	}
	defer db.Close()

	// UseCase
	u := usecase.NewGetMessage(messageID, repository.NewMessageRepository(db))
	if message, err = u.GetMessageUseCase(r.Context()); err != nil {
		if err == usecase.ErrNotExistsData {
			err = helper.NewNotFoundError(err.Error())
			return
		}
		err = helper.NewInternalServerError(err.Error())
		return
	}
	return
}

func controlDeleteMessage(r *http.Request) error {
	// get parameter
	vars := mux.Vars(r)
	paramMessageID, ok := vars["id"]
	if !ok {
		return helper.NewBadRequestError("id is required")
	}
	messageID, err := strconv.Atoi(paramMessageID)
	if err != nil {
		return helper.NewInternalServerError(err.Error())
	}

	// db connect
	db, err := mysql.NewDB()
	if err != nil {
		return helper.NewInternalServerError(err.Error())
	}
	defer db.Close()

	// UseCase
	u := usecase.NewDeleteMessage(messageID, repository.NewMessageRepository(db))
	if err = u.DeleteMessageUseCase(r.Context()); err != nil {
		return helper.NewInternalServerError(err.Error())
	}
	return err
}
