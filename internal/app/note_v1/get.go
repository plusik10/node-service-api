package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode)
	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := squirrel.Select("author,text,title").
		From(noteTable).
		Where(squirrel.Eq{"id": req.GetId()}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var text string
	var author string
	var title string

	row.Next()
	err = row.Scan(&author, &text, &title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("Record with id %s doest not exist", req.Id)
		}
		return nil, err
	}

	log.Println("a request to retrieve a note has been completed")
	return &desc.GetResponse{Title: title, Author: author, Text: text}, nil
}
