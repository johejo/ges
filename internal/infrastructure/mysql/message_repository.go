package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/huandu/go-sqlbuilder"

	"github.com/johejo/ges/internal/domain/model"
	"github.com/johejo/ges/internal/domain/repository"
)

const messageTable = "message"

var messageTableColumns = []string{"id", "title", "text"}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repository.MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) Save(ctx context.Context, m *model.Message) error {
	mj, err := m.ToJSON()
	if err != nil {
		return err
	}

	sb := sqlbuilder.NewInsertBuilder()
	_sql, args := sb.InsertInto(messageTable).Cols(messageTableColumns...).Values(mj.ID, mj.Title, mj.Text).Build()
	logger.Println(_sql, args)

	result, err := r.db.ExecContext(ctx, _sql, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("rows != 1")
	}

	return nil
}

func (r *messageRepository) Load(ctx context.Context, id string) (*model.Message, error) {
	sb := sqlbuilder.NewSelectBuilder()
	_sql, args := sb.Select(messageTableColumns...).From(messageTable).Where(sb.Equal("id", id)).Build()
	logger.Println(_sql, args)

	var m model.MessageJSON
	if err := r.db.QueryRowContext(ctx, _sql, args...).Scan(&m.ID, &m.Title, &m.Text); err != nil {
		return nil, err
	}

	return m.ToMessage()
}

func (r *messageRepository) LoadAll(ctx context.Context) ([]*model.Message, error) {
	sb := sqlbuilder.NewSelectBuilder()
	_sql, args := sb.Select(messageTableColumns...).From(messageTable).Build()
	logger.Println(_sql, args)

	rows, err := r.db.QueryContext(ctx, _sql, args...)
	if err != nil {
		return nil, err
	}

	var mj model.MessageJSON
	ms := make([]*model.Message, 0)
	for rows.Next() {
		if err := rows.Scan(&mj.ID, &mj.Title, &mj.Text); err != nil {
			return nil, err
		}
		m, err := mj.ToMessage()
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil
}
