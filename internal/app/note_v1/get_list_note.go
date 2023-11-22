package note_v1

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) GetListNote(ctx context.Context, empty *desc.Empty) (*desc.GetListNoteResponse, error) {

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

	return &desc.GetListNoteResponse{Notes: notes}, nil

}
