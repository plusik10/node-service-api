package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/plusik10/note-service-api/internal/repository/const"
	"github.com/plusik10/note-service-api/internal/service/note"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"github.com/plusik10/note-service-api/pkg/postgres"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ note.INoteRepository = (*PostgresRepository)(nil)

type PostgresRepository struct {
	pg *postgres.Postgres
}

func (p *PostgresRepository) Get(ctx context.Context, id int64) (desc.Note, error) {
	note := desc.Note{Id: id}
	query, arg, err := p.pg.Builder.
		Select(repository.Author, repository.Title, repository.Text, repository.CreatedAt, repository.UpdatedAt).
		From(repository.NoteTable).
		Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		fmt.Printf("postgresRepository - pg.Builder - err: %s", err.Error())
		return desc.Note{}, fmt.Errorf("pgRepository - pg.Builder: %s", err.Error())
	}

	row, err := p.pg.Db.QueryContext(ctx, query, arg...)
	if err != nil {
		fmt.Printf("postgresRepository - get - queryContext err: %s", err.Error())
		return desc.Note{}, fmt.Errorf("repository - get - QueryContext err: %s", err.Error())
	}
	defer row.Close()

	var (
		CreateAt time.Time
		UpdateAt sql.NullTime
	)

	row.Next()
	err = row.Scan(&note.Author, &note.Title, &note.Text, &CreateAt, &UpdateAt)
	if err != nil {
		fmt.Printf("postgresRepository - get - Scan err: %s", err.Error())
		return desc.Note{}, fmt.Errorf("repository - get - Scan. Err: %s", err.Error())
	}

	note.CreatedAt = timestamppb.New(CreateAt)
	if UpdateAt.Valid {
		note.UpdatedAt = timestamppb.New(UpdateAt.Time)
	} else {
		note.UpdatedAt = nil
	}

	return note, nil
}

func (p *PostgresRepository) GetAll(ctx context.Context) ([]*desc.Note, error) {
	query, arg, err := p.pg.Builder.
		Select(repository.Id,
			repository.Author,
			repository.Title,
			repository.Text,
			repository.UpdatedAt,
			repository.CreatedAt).
		From(repository.NoteTable).ToSql()
	if err != nil {
		fmt.Printf("repository - getAll - Builder err: %s\n", err.Error())
		return nil, fmt.Errorf("postgresRepository - getall - builder err: %s", err.Error())
	}

	rows, err := p.pg.Db.QueryContext(ctx, query, arg...)
	if err != nil {

		fmt.Printf("repository - getall - queryContext err: %s\n", err.Error())
		return nil, fmt.Errorf("postgresRepository - getall - queryContext err:%s", err.Error())
	}
	defer rows.Close()

	notes := []*desc.Note{}
	var (
		CreateAt time.Time
		UpdateAt sql.NullTime
	)

	for rows.Next() {
		note := &desc.Note{}
		errScan := rows.Scan(&note.Id, &note.Author, &note.Title, &note.Text, &UpdateAt, &CreateAt)
		if errScan != nil {
			fmt.Printf("repository - getall - scan err: %s", err.Error())
			return nil, err
		}
		note.CreatedAt = timestamppb.New(CreateAt)
		if UpdateAt.Valid {
			note.UpdatedAt = timestamppb.New(UpdateAt.Time)
		} else {
			note.UpdatedAt = nil
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (p *PostgresRepository) Update(ctx context.Context, id int64, title string, author string, text string) error {
	query, arg, err := p.pg.Builder.
		Update(repository.NoteTable).
		Set(repository.Author, author).
		Set(repository.Title, title).
		Set(repository.Text, text).
		Set(repository.UpdatedAt, "now()").
		Where(squirrel.Eq{repository.Id: id}).
		ToSql()

	_, err = p.pg.Db.ExecContext(ctx, query, arg...)
	if err != nil {
		return err
	}
	fmt.Printf("note with id=%d was deleted\n", id)

	return nil
}

func (p *PostgresRepository) Delete(ctx context.Context, id int64) error {
	query, arg, err := p.pg.Builder.Delete(repository.NoteTable).Where(squirrel.Eq{"id": id}).ToSql()
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
	query, arg, err := p.pg.Builder.Insert(repository.NoteTable).
		Columns(repository.Author, repository.Title, repository.Text).
		Values(author, title, text).Suffix("returning id").ToSql()
	if err != nil {
		return 0, fmt.Errorf("repository - Create - Builder err: %s", err.Error())
	}

	row, err := p.pg.Db.QueryContext(ctx, query, arg...)
	if err != nil {
		return 0, fmt.Errorf("repository - Create - QueryContext err: %s", err.Error())
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository - create - Scan. Err: %s", err.Error())
	}

	return id, nil
}

func NewPostgresRepository(postgres *postgres.Postgres) *PostgresRepository {
	return &PostgresRepository{pg: postgres}
}
