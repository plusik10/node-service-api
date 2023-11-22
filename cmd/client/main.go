package main

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"log"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}

	defer conn.Close()

	client := desc.NewNoteV1Client(conn)

	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "First Request GRPC!",
		Author: "Konstantin",
		Text:   "Дорогой дневник, мне не подобрать слов что бы описать боль и унижение которые я испытал сегодня...",
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Result:", res.String())

}
