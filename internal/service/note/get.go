package note

import (
	"context"

	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (s *Service) Get(ctx context.Context, id int64) (desc.Note, error) {
	note, err := s.rep.Get(ctx, id)
	if err != nil {
		return desc.Note{}, err
	}

	return note, nil
}
