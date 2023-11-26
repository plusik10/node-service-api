package note_v1

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, empty *emptypb.Empty) (*desc.GetListResponse, error) {
	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := squirrel.Select(colAuthor, colText, colAuthor).
		From(noteTable).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var author string
	var text string
	var title string

	notes := []*desc.Note{}
	for rows.Next() {
		err := rows.Scan(&author, &text, &title)
		if err != nil {
			return nil, err
		}
		note := &desc.Note{Author: author, Title: title, Text: text}
		notes = append(notes, note)
	}

	log.Println("a request to retrieve a notes has been completed")
	return &desc.GetListResponse{Notes: notes}, nil
}
