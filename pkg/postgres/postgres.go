package postgres

import (
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Db      *sqlx.DB
}

func New(url string) (*Postgres, error) {
	fmt.Println("Ya tut")
	fmt.Println(url)
	db, err := sqlx.Open("pgx", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error connecting to the database: %s \n", err)
		return nil, err //TODO Завернуть в ошибку
	}

	pg := &Postgres{Db: db}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	log.Printf("data base connection successful\n")

	return pg, nil
}

func (p *Postgres) CloseConnect() error {
	if err := p.Db.Close(); err != nil {
		return err
	}
	return nil
}
