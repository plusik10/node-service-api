package service

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

type NoteService interface {
	Create(ctx context.Context, title string, author string, text string) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (model.Note, error)
	GetAll(ctx context.Context) ([]model.Note, error)
	Update(ctx context.Context, id int64, title string, author string, text string) error
}
