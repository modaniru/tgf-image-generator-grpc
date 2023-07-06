package main

import (
	"log"

	"github.com/modaniru/image-generator/src/client"
	"github.com/modaniru/image-generator/src/service"
	"google.golang.org/grpc"
)

// Проблема вставки шрифта  в картинку
// Проблема долгой загрузки РЕШЕНО

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	tgf := client.NewTgfClient(conn)
	service := service.NewService(tgf)
	_, err = service.GenerateImage([]string{"snivanov"})
	if err != nil {
		log.Fatal(err.Error())
	}
}
