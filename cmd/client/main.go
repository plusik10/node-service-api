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

	fmt.Println("\n CreateNote")
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "First Request GRPC!",
		Author: "Konstantin",
		Text:   "Дорогой дневник, мне не подобрать слов что бы описать боль и унижение которые я испытал сегодня...",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", res.String())

	fmt.Println("\n DeleteNote")
	delres, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", delres.String())

	fmt.Println("\n GetListNote")
	lists, err := client.GetListNote(context.Background(), &desc.Empty{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", lists.String())

	fmt.Println("\n GetNote")
	get_note, err := client.GetNote(context.Background(), &desc.GetNoteRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", get_note.String())

	fmt.Println("\n UpdateNote")
	updNote, err := client.UpdateNote(context.Background(), &desc.Note{Id: 1, Author: "Max", Text: "Hello world", Title: "Add"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", updNote.String())

}
