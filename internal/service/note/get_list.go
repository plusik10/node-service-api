package note

import (
	"context"
	"fmt"

	"github.com/plusik10/note-service-api/internal/model"
)

func (s *service) GetAll(ctx context.Context) ([]model.Note, error) {
	notes, err := s.repo.GetAll(ctx)
	if err != nil {
		fmt.Printf("note.Service - GetALL - Repository: %s", err.Error())
		return nil, err
	}

	return notes, nil
}
