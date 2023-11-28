package note_v1

import (
	"context"
	"fmt"
	"log"

	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, req *emptypb.Empty) (*desc.GetListResponse, error) {
	notes, err := n.service.GetAll(ctx)
	if err != nil {
		fmt.Printf("note_v1 - GetList - service  err:%s", err.Error())
		return nil, err
	}
	log.Println("a request to retrieve a note has been completed")
	return &desc.GetListResponse{Notes: notes}, nil
}
