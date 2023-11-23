package main

import (
	"context"
	"fmt"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}

	defer conn.Close()

	Note := desc.NewNoteV1Client(conn)

	fmt.Println("\n CreateNote")
	res, err := Note.Create(context.Background(), &desc.CreateRequest{
		Title:  "First Request GRPC!",
		Author: "Konstantin",
		Text:   "Дорогой дневник, мне не подобрать слов что бы описать боль и унижение которые я испытал сегодня...",
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", res.String())

	fmt.Println("\n DeleteNote")
	delres, err := Note.Delete(context.Background(), &desc.DeleteRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", delres.String())

	fmt.Println("\n GetListNote")
	lists, err := Note.GetList(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", lists.String())

	fmt.Println("\n GetNote")
	get_note, err := Note.Get(context.Background(), &desc.GetRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", get_note.String())

	fmt.Println("\n UpdateNote")
	updNote, err := Note.Update(context.Background(),
		&desc.UpdateRequest{Note: &desc.Note{Id: 1, Author: "Max", Text: "Hello world", Title: "Add"}})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Result:", updNote.String())

}
