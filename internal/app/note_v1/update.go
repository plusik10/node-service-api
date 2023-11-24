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

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query, arg, err := squirrel.Update(noteTable).
		PlaceholderFormat(squirrel.Dollar).
		Set("author", req.Note.Author).
		Set("title", req.Note.Title).
		Set("text", req.Note.GetText()).
		Set("updated_at", "now()").
		Where(squirrel.Eq{"id": req.Note.Id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.QueryContext(ctx, query, arg...)
	if err != nil {
		return nil, err
	}

	log.Printf("note with id = %d was updated", req.Note.GetId())
	return &emptypb.Empty{}, nil
}
