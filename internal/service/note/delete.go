package note

import (
	"context"
	"fmt"
)

func (s *service) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting a record: %s", err.Error())
	}

	return nil
}
