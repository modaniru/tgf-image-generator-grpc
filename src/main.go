package main

import (
	"image/color"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/modaniru/image-generator/src/client"
	"github.com/modaniru/image-generator/src/draw"
	"github.com/modaniru/image-generator/src/service"
	"github.com/modaniru/image-generator/src/utils"
	"google.golang.org/grpc"
)

// Проблема вставки шрифта  в картинку
// Проблема долгой загрузки РЕШЕНО
// TODO documentation in README.md
// TODO viper нужен  ли, можно же все в .env файлы ВРОДЕ можно скипать и все писать в .env
// TODO .env file
// TODO нормальная галочка, если человек партнер, серая, если компаньен
// TODO create normal base style
func main() {
	_ = godotenv.Load()
	images, err := utils.LoadPngImages("images")
	if err != nil {
		log.Fatal(err.Error())
	}
	// TODO security
	conn, err := grpc.Dial(os.Getenv("TGF_SERVICE_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	tgf := client.NewTgfClient(conn)
	httpClient := client.NewHttpClient(&http.Client{
		Timeout: 5 * time.Second,
	})
	s := service.NewDrawerService(images, &draw.ImageStyle{
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
	})
	service := service.NewService(tgf, httpClient, s)
	_, err = service.GenerateImage([]string{"modaniru", "snivanov"})
	if err != nil {
		log.Fatal(err.Error())
	}
}
