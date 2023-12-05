package note

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/plusik10/note-service-api/internal/model"
	"github.com/plusik10/note-service-api/internal/repository"
	constantRepository "github.com/plusik10/note-service-api/internal/repository/note/const"
	"github.com/plusik10/note-service-api/pkg/postgres"
)

var _ repository.NoteRepository = (*PostgresRepository)(nil)

type PostgresRepository struct {
	pg *postgres.Postgres
}

func (p *PostgresRepository) Get(ctx context.Context, id int64) (model.Note, error) {
	note := model.Note{Id: id}
	query, arg, err := p.pg.Builder.
		Select(constantRepository.Author,
			constantRepository.Title,
			constantRepository.Text,
			constantRepository.CreatedAt,
			constantRepository.UpdatedAt).
		From(constantRepository.NoteTable).
		Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return model.Note{}, err
	}

	row, err := p.pg.Db.QueryContext(ctx, query, arg...)
	if err != nil {
		return model.Note{}, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&note.Author, &note.Title, &note.Text, &note.CreateAt, &note.UpdateAt)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (p *PostgresRepository) GetAll(ctx context.Context) ([]model.Note, error) {
	query, arg, err := p.pg.Builder.
		Select(constantRepository.Id,
			constantRepository.Author,
			constantRepository.Title,
			constantRepository.Text,
			constantRepository.UpdatedAt,
			constantRepository.CreatedAt).
		From(constantRepository.NoteTable).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.pg.Db.QueryContext(ctx, query, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notes := []model.Note{}

	for rows.Next() {
		note := model.Note{}
		err = rows.Scan(&note.Id, &note.Author, &note.Title, &note.Text, &note.UpdateAt, &note.CreateAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (p *PostgresRepository) Update(ctx context.Context, id int64, title string, author string, text string) error {
	query, arg, err := p.pg.Builder.
		Update(constantRepository.NoteTable).
		Set(constantRepository.Author, author).
		Set(constantRepository.Title, title).
		Set(constantRepository.Text, text).
		Set(constantRepository.UpdatedAt, "now()").
		Where(squirrel.Eq{constantRepository.Id: id}).
		ToSql()

	_, err = p.pg.Db.ExecContext(ctx, query, arg...)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) Delete(ctx context.Context, id int64) error {
	query, arg, err := p.pg.Builder.Delete(constantRepository.NoteTable).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = p.pg.Db.ExecContext(ctx, query, arg...)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) Create(ctx context.Context, author string, title string, text string) (int64, error) {
	query, arg, err := p.pg.Builder.
		Insert(constantRepository.NoteTable).
		Columns(constantRepository.Author,
			constantRepository.Title,
			constantRepository.Text).
		Values(author, title, text).Suffix("returning id").ToSql()
	if err != nil {
		return 0, err
	}

	row, err := p.pg.Db.QueryContext(ctx, query, arg...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func NewPostgresRepository(postgres *postgres.Postgres) *PostgresRepository {
	return &PostgresRepository{pg: postgres}
}
