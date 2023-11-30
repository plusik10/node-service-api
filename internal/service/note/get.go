package note

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

func (s *service) Get(ctx context.Context, id int64) (model.Note, error) {
	note, err := s.repo.Get(ctx, id)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}
