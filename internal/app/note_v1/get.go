package note_v1

import (
	"context"
	"log"

	"github.com/plusik10/note-service-api/internal/converter"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := n.noteService.Get(ctx, req.GetId())
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &desc.GetResponse{Note: converter.ToDescFromNote(note)}, nil
}
