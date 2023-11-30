package note

import (
	"github.com/plusik10/note-service-api/internal/repository"
	def "github.com/plusik10/note-service-api/internal/service"
)

var _ def.NoteService = (*service)(nil)

type service struct {
	repo repository.NoteRepository
}

func NewService(repo repository.NoteRepository) def.NoteService {
	return &service{repo: repo}
}
