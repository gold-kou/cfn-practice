package usecase

import (
	"context"

	modelHTTP "github.com/gold-kou/cfn-practice/backend/app/adapter/http/model"
	"github.com/gold-kou/cfn-practice/backend/app/domain/model"
	"github.com/gold-kou/cfn-practice/backend/app/domain/repository"
)

type RegisterMessageUseCaseInterface interface {
	RegisterMessageUseCase(context.Context) error
}

type RegisterMessage struct {
	paramRegisterMessage modelHTTP.RequestRegisterMessage
	messageRepo          repository.MessageRepository
}

func NewRegisterMessage(reqRegisterMessage modelHTTP.RequestRegisterMessage, messageRepo repository.MessageRepository) *RegisterMessage {
	return &RegisterMessage{
		paramRegisterMessage: reqRegisterMessage,
		messageRepo:          messageRepo,
	}
}

func (message *RegisterMessage) RegisterMessageUseCase(ctx context.Context) error {
	m := model.Message{
		Message: message.paramRegisterMessage.Message,
	}
	err := message.messageRepo.Create(ctx, m)
	if err != nil {
		if err == repository.ErrDuplicateData {
			return ErrDuplicateData
		}
		return err
	}
	return nil
}
