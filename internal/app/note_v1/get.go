package note_v1

import (
	"context"
	"log"

	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := n.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	log.Printf("a request to retrieve a note has been completed id=%d \n", req.GetId())
	return &desc.GetResponse{Note: &note}, nil
}
