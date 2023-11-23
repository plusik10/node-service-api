package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, empty *emptypb.Empty) (*desc.GetListResponse, error) {

	fmt.Println("Success! GetListNote")
	notes := []*desc.Note{
		&desc.Note{
			Author: "Konstantin",
			Title:  "А что если это все непанастаясему",
			Text:   "fis",
		},
		&desc.Note{
			Author: "Misha",
			Title:  "Новость",
			Text:   "Все очень плохо",
		},
		&desc.Note{
			Author: "Oleg",
			Title:  "News",
			Text:   "Да что ты ноешь",
		},
	}

	return &desc.GetListResponse{Notes: notes}, nil

}
