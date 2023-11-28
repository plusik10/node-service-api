package note

import (
	"context"
	"fmt"

	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (s *Service) GetAll(ctx context.Context) ([]*desc.Note, error) {
	notes, err := s.rep.GetAll(ctx)
	if err != nil {
		fmt.Printf("note.Service - GetALL - Repository: %s", err.Error())
		return nil, err
	}

	return notes, nil
}
