package note_v1

import (
	"context"

	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := n.noteService.Update(
		ctx,
		req.GetId(),
		req.GetTitle(),
		req.GetAuthor(),
		req.GetText(),
	)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
