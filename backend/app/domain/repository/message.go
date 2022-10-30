package repository

import (
	"context"

	"github.com/gold-kou/cfn-practice/backend/app/domain/model"
)

type MessageRepository interface {
	Create(ctx context.Context, message model.Message) error
	Find(ctx context.Context, id int) (model.Message, error)
	Remove(ctx context.Context, id int) error
}
