package note

import (
	"context"

	"github.com/plusik10/note-service-api/internal/app/note_v1"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

var _ note_v1.NoteService = (*Service)(nil)

type INoteRepository interface {
	Create(ctx context.Context, author string, title string, text string) (int64, error)
	Get(ctx context.Context, id int64) (desc.Note, error)
	GetAll(ctx context.Context) ([]*desc.Note, error)
	Update(ctx context.Context, id int64, title string, author string, text string) error
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	rep INoteRepository
}

func New(rep INoteRepository) *Service {
	return &Service{rep: rep}
}
