package main

import (
	"fmt"
	"github.com/plusik10/note-service-api/internal/app/note_v1"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to mapping port: %s", err.Error())
	}

	s := grpc.NewServer()
	desc.RegisterNoteV1Server(s, note_v1.NewNote())

	fmt.Println("Server is running")
	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}

}
