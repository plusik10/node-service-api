package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plusik10/note-service-api/config"
	"github.com/plusik10/note-service-api/internal/app/note_v1"
	"github.com/plusik10/note-service-api/internal/repository"
	"github.com/plusik10/note-service-api/internal/service/note"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"github.com/plusik10/note-service-api/pkg/postgres"
	"google.golang.org/grpc"
)

func Run(cfg *config.Config) error {
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		return fmt.Errorf("app - run - initPostgres err: %w", err)
	}
	defer pg.CloseConnect()

	pgRepo := repository.NewPostgresRepository(pg)
	noteService := note.New(pgRepo)
	note := note_v1.NewNote(noteService)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		startHTTP(cfg)
	}()
	go func() {
		defer wg.Done()
		startGRPC(cfg.GRPC.Port, note)
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func startHTTP(cfg *config.Config) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, mux, cfg.GRPC.Port, opts)
	if err != nil {
		log.Fatalf("failed to mapping port: %s \n ", cfg.HTTP.Port)
		return err
	}
	fmt.Printf("HTTPServer is running on : %s\n", cfg.HTTP.Port)

	return http.ListenAndServe(cfg.HTTP.Port, mux)
}

func startGRPC(grpcPort string, server desc.NoteV1Server) error {
	list, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()))
	desc.RegisterNoteV1Server(s, server)
	fmt.Printf("GRPCServer is running on : %s\n", grpcPort)

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}
