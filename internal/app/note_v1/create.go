package note_v1

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbUser     = "postgres"
	dbPassword = "qwerty"
	dbName     = "note-service"
	sslMode    = "disable"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, arg, err := squirrel.Insert(noteTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns("author,title,text").
		Values(req.GetAuthor(), req.GetTitle(), req.GetText()).
		Suffix("returning id").ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, arg...)
	if err != nil {
		return nil, err
	}

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	log.Printf("add new note id: %d\n", id)
	return &desc.CreateResponse{
		Id: id,
	}, nil
}
