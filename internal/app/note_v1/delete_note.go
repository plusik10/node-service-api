package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, request *desc.DeleteNoteRequest) (*desc.Empty, error) {
	fmt.Printf("Success! object  %s deleted", request.GetId())
	return &desc.Empty{}, nil
}
