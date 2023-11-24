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

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	dbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, arg, err := squirrel.Delete(noteTable).
		Where(squirrel.Eq{"id": req.Id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.QueryContext(ctx, query, arg...)
	if err != nil {
		return nil, err
	}

	log.Printf("note with id %d was deleted \n", req.GetId())
	return &emptypb.Empty{}, nil
}
