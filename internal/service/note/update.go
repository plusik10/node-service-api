package note

import (
	"context"
	"fmt"
)

func (s *Service) Update(ctx context.Context, id int64, title string, author string, text string) error {
	fmt.Println("I'm Here!")
	if err := s.rep.Update(ctx, id, title, author, text); err != nil {
		return fmt.Errorf("error updating a record: %s", err.Error())
	}

	return nil
}
