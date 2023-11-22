package note_v1

import desc "github.com/plusik10/note-service-api/pkg/note_v1"

var _ desc.NoteV1Server = (*Note)(nil)

type Note struct {
	desc.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}
