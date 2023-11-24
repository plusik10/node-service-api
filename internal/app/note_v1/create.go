package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, cr *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println("Success!")
	fmt.Printf("title: %s", cr.GetTitle())
	fmt.Printf("Author: %s", cr.GetAuthor())
	fmt.Printf("text: %s", cr.GetText())

	return &desc.CreateResponse{
		Id: 1,
	}, nil
}
