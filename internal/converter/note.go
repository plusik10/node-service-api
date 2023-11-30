package converter

import (
	"github.com/plusik10/note-service-api/internal/model"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToDescFromNotes(notes []model.Note) []*desc.Note {
	resultNotes := make([]*desc.Note, 0, len(notes))
	var (
		updatedAt *timestamppb.Timestamp
		createdAt *timestamppb.Timestamp
	)
	for _, n := range notes {
		createdAt = timestamppb.New(n.CreateAt)
		if n.UpdateAt != nil {
			updatedAt = timestamppb.New(*n.UpdateAt)
		}
		resultNotes = append(resultNotes, &desc.Note{Id: n.Id,
			Title:     n.Title,
			Author:    n.Author,
			Text:      n.Text,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return resultNotes
}

func ToDescFromNote(note model.Note) *desc.Note {
	var (
		updatedAt *timestamppb.Timestamp
		createdAt *timestamppb.Timestamp
	)
	if note.UpdateAt != nil {
		updatedAt = timestamppb.New(*note.UpdateAt)
	}
	createdAt = timestamppb.New(note.CreateAt)

	return &desc.Note{
		Id:        note.Id,
		Title:     note.Title,
		Text:      note.Text,
		Author:    note.Author,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
}
