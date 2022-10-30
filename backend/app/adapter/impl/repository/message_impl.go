package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"

	mysqlAdapter "github.com/gold-kou/cfn-practice/backend/app/adapter/mysql"
	"github.com/gold-kou/cfn-practice/backend/app/domain/model"
	"github.com/gold-kou/cfn-practice/backend/app/domain/repository"
)

type messageRepositoryImpl struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repository.MessageRepository {
	return &messageRepositoryImpl{db}
}

func (r *messageRepositoryImpl) Create(ctx context.Context, message model.Message) (err error) {
	q := "INSERT INTO `messages` (`message`) VALUES (?)"
	tx := mysqlAdapter.GetTransaction(ctx)
	if tx != nil {
		_, err = tx.ExecContext(ctx, q, message.Message)
	} else {
		_, err = r.db.ExecContext(ctx, q, message.Message)
	}
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == mysqlAdapter.DuplicateEntryErrorNumber {
			return xerrors.Errorf("%w", repository.ErrDuplicateData)
		}
		return xerrors.Errorf("insert into messages: %v", err)
	}
	return
}

func (r *messageRepositoryImpl) Find(ctx context.Context, id int) (message model.Message, err error) {
	q := "SELECT `id`, `message` FROM `messages` WHERE `id` = ?"
	err = r.db.QueryRowContext(ctx, q, id).Scan(&message.ID, &message.Message)
	if err == sql.ErrNoRows {
		err = repository.ErrNotExistsData
		return
	}
	return
}

func (r *messageRepositoryImpl) Remove(ctx context.Context, id int) (err error) {
	q := "DELETE FROM `messages` WHERE `id` = ?"
	tx := mysqlAdapter.GetTransaction(ctx)
	if tx != nil {
		_, err = tx.ExecContext(ctx, q, id)
	} else {
		_, err = r.db.ExecContext(ctx, q, id)
	}
	return
}
