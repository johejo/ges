package usecase

import (
	"context"

	"github.com/johejo/ges/internal/domain/model"
	"github.com/johejo/ges/internal/domain/repository"
)

type MessageUseCase interface {
	Save(ctx context.Context, m *model.Message) error
	Load(ctx context.Context, id string) (*model.Message, error)
	LoadAll(ctx context.Context) ([]*model.Message, error)
}

type messageUseCase struct {
	repository.MessageRepository
}

func NewMessageUseCase(mr repository.MessageRepository) MessageUseCase {
	return &messageUseCase{
		MessageRepository: mr,
	}
}
