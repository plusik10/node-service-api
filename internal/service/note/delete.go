package note

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, id int64) error {
	if err := s.rep.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting a record: %s", err.Error())
	}

	return nil
}
