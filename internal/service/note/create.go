package note

import "context"

func (s *Service) Create(ctx context.Context, title string, author string, text string) (int64, error) {
	id, err := s.rep.Create(ctx, title, author, text)
	if err != nil {
		return 0, err
	}

	return id, nil
}
