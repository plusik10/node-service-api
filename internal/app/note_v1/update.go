package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Note is update")
	fmt.Println(req.Note.GetId())
	fmt.Println(req.Note.GetTitle())
	fmt.Println(req.Note.GetAuthor())
	fmt.Println(req.Note.GetText())
	return &emptypb.Empty{}, nil
}
