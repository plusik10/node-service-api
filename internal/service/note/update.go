package note

import (
	"context"
	"fmt"
)

func (s *service) Update(ctx context.Context, id int64, title string, author string, text string) error {
	if err := s.repo.Update(ctx, id, title, author, text); err != nil {
		return fmt.Errorf("error updating a record: %s", err.Error())
	}

	return nil
}
