package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, request *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("Success! GetNote")
	return &desc.GetNoteResponse{
		Author: "Konstantin",
		Title:  "GRPC",
		Text:   "Как тебе такой грпс илон маск?",
	}, nil
}
