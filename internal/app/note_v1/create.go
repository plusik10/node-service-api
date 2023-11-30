package note_v1

import (
	"context"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := n.noteService.Create(ctx, req.GetTitle(), req.GetAuthor(), req.GetText())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
