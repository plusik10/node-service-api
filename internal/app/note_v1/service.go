package note_v1

import (
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

type Note struct {
	desc.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}
