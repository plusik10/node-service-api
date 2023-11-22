package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, note *desc.Note) (*desc.Empty, error) {
	fmt.Println("Note is update")
	fmt.Println(note.GetId())
	fmt.Println(note.GetTitle())
	fmt.Println(note.GetAuthor())
	fmt.Println(note.GetText())
	return &desc.Empty{}, nil
}
