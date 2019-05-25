package repository

import (
	"context"

	"github.com/johejo/ges/internal/domain/model"
)

type MessageRepository interface {
	Save(ctx context.Context, m *model.Message) error
	Load(ctx context.Context, id string) (*model.Message, error)
	LoadAll(ctx context.Context) ([]*model.Message, error)
}
