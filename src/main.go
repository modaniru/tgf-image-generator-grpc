package main

import (
	"image/color"
	"log"
	"net"
	"net/http"
	"os"

	pkg "github.com/modaniru/image-generator/pkg/proto"
	"github.com/modaniru/image-generator/src/client"
	"github.com/modaniru/image-generator/src/draw"
	"github.com/modaniru/image-generator/src/server"
	"github.com/modaniru/image-generator/src/service"
	"github.com/modaniru/image-generator/src/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO file cross platform
func main() {
	// Load .env file if exists
	utils.LoadConfig()
	// Load static images
	staticImages, err := utils.LoadPngImages("images")
	if err != nil {
		log.Fatal(err.Error())
	}
	// Dependency Injection
	httpClient := client.NewHttpClient(&http.Client{})
	tgfClientConnection, err := grpc.Dial(os.Getenv("TGF_SERVICE_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer tgfClientConnection.Close()
	tgfClient := client.NewTgfClient(tgfClientConnection)
	drawerService := service.NewDrawerService(staticImages, getImageStyle())
	imageGeneratorService := service.NewImageGeneratorService(tgfClient, httpClient, drawerService)
	ImageServer := server.NewImageServiceServer(imageGeneratorService)

	// create listener
	listener, err := net.Listen("tcp", "localhost:" + utils.GetPort())
	if err != nil {
		log.Fatal(err.Error())
	}
	// create server
	grpcServer := grpc.NewServer()
	pkg.RegisterImageServiceServer(grpcServer, ImageServer)
	// reflection
	reflection.Register(grpcServer)
	grpcServer.Serve(listener)
}

// default image style
func getImageStyle() *draw.ImageStyle{
	return &draw.ImageStyle{
		Background: []color.Color{color.RGBA{196, 113, 242, 255}, color.White, color.RGBA{247, 108, 198, 255}},
		HeaderStyle: draw.HeaderStyle{
			Background: []color.Color{color.RGBA{92, 115, 185, 255}, color.RGBA{179, 48, 225, 255}},
			TextColors: color.White,
		},
		GeneralFollowsCard: draw.GeneralFollowsCard{
			Background:      []color.Color{color.RGBA{235, 244, 245, 255}, color.RGBA{181, 198, 224, 255}},
			TextColors:      color.Black,
			ImageBackground: []color.Color{color.White},
		},
		InputedUsersCard: draw.InputedUsersCard{
			Background:      []color.Color{color.RGBA{235, 244, 245, 255}, color.RGBA{181, 198, 224, 255}},
			TextColors:      color.Black,
			ImageBackground: []color.Color{color.White},
		},
	}
}