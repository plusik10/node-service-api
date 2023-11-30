package note

import "context"

func (s *service) Create(ctx context.Context, title string, author string, text string) (int64, error) {
	id, err := s.repo.Create(ctx, title, author, text)
	if err != nil {
		return 0, err
	}

	return id, nil
}
