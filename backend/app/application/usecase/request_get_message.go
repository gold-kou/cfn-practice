package usecase

import (
	"context"

	"github.com/gold-kou/cfn-practice/backend/app/domain/model"
	"github.com/gold-kou/cfn-practice/backend/app/domain/repository"
)

type GetMessageUseCaseInterface interface {
	GetMessageUseCase(ctx context.Context, id int) (model.Message, error)
}

type GetMessage struct {
	id          int
	messageRepo repository.MessageRepository
}

func NewGetMessage(id int, messageRepo repository.MessageRepository) *GetMessage {
	return &GetMessage{
		id:          id,
		messageRepo: messageRepo,
	}
}

func (message *GetMessage) GetMessageUseCase(ctx context.Context) (m model.Message, err error) {
	m, err = message.messageRepo.Find(ctx, message.id)
	if err != nil {
		if err == repository.ErrNotExistsData {
			return m, ErrNotExistsData
		}
		return
	}
	return
}
