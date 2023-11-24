package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Printf("Success! object  %d deleted", request.GetId())
	return &emptypb.Empty{}, nil
}
