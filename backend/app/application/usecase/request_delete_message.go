package usecase

import (
	"context"

	"github.com/gold-kou/cfn-practice/backend/app/domain/repository"
)

type DeleteMessageUseCaseInterface interface {
	DeleteMessageUseCase(ctx context.Context, id int) error
}

type DeleteMessage struct {
	id          int
	messageRepo repository.MessageRepository
}

func NewDeleteMessage(id int, messageRepo repository.MessageRepository) *DeleteMessage {
	return &DeleteMessage{
		id:          id,
		messageRepo: messageRepo,
	}
}

func (message *DeleteMessage) DeleteMessageUseCase(ctx context.Context) error {
	err := message.messageRepo.Remove(ctx, message.id)
	if err != nil {
		return err
	}
	return nil
}
