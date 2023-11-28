package note_v1

import (
	"context"

	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

var _ desc.NoteV1Server = (*Note)(nil)

const (
	noteTable     = "note"
	host          = "localhost"
	port          = "54321"
	dbUser        = "postgres"
	dbPassword    = "qwerty"
	dbName        = "note-service"
	sslMode       = "disable"
	colAuthor     = "author"
	colTitle      = "title"
	colText       = "text"
	colUpdated_at = "updated_at"
)

type NoteService interface {
	Create(ctx context.Context, title string, author string, text string) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (desc.Note, error)
	GetAll(ctx context.Context) ([]*desc.Note, error)
	Update(ctx context.Context, id int64, title string, author string, text string) error
}

type Note struct {
	desc.UnimplementedNoteV1Server
	service NoteService
}

func NewNote(service NoteService) *Note {
	return &Note{service: service}
}
