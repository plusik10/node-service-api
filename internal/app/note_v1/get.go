package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Success! GetNote")
	return &desc.GetResponse{
		Author: "Konstantin",
		Title:  "GRPC",
		Text:   "Как тебе такой грпс илон маск?",
	}, nil
}
