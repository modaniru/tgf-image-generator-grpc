package service

import (
	"context"
	"image"
	"image/color"

	"github.com/fogleman/gg"
	pkg "github.com/modaniru/image-generator/pkg/proto"
	"github.com/modaniru/image-generator/src/client"
	"github.com/modaniru/image-generator/src/utils"
)

type Service struct {
	tgfClient *client.TgfClient
}

func NewService(tgfClient *client.TgfClient) *Service {
	return &Service{tgfClient: tgfClient}
}

func (s *Service) GenerateImage(nicknames []string) (image.Image, error) {
	res, err := s.tgfClient.GetGeneralFollows(context.Background(), &pkg.GetTGFRequest{Usernames: nicknames})
	if err != nil {
		return nil, err
	}
	Height := 240.0
	Width := 1120.0

	followsCount := len(res.GeneralStreamers)

	rowsCount := followsCount / 5
	if rowsCount*5 < followsCount {
		rowsCount++
	}
	inputCount := len(res.InputedUsers)
	inputedUsersRows := inputCount / 5
	if inputedUsersRows*5 < inputCount {
		inputedUsersRows++
	}
	inputUserCardHeight := 100.0
	cardHeight := 240.0
	iHeight := inputUserCardHeight * float64(inputedUsersRows)
	Height += float64(rowsCount)*cardHeight + iHeight

	c := gg.NewContext(int(Width), int(Height))

	utils.FillWithGradient(c, 0, 0, Width, Height, color.RGBA{238, 130, 238, 255}, color.RGBA{211, 247, 253, 255}, color.RGBA{0, 209, 255, 255})

	err = utils.DrawHeaderSquareWithShadow(c, 20, 20, float64(Width)-40, 200, "Twitch General Follows")
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]*pkg.ResponseUser)
	for _, v := range res.InputedUsers {
		userMap[v.DisplayName] = v
	}

	imageMap := make(map[string]image.Image)

	channel := make(chan *utils.Response)
	for _, v := range res.GeneralStreamers {
		go utils.GetImageFromURI(v.Streamer.ImageLink, v.Streamer.DisplayName, channel)
	}
	for _, v := range res.InputedUsers {
		go utils.GetImageFromURI(v.ImageLink, v.DisplayName, channel)
	}
	for i := 0; i < len(res.GeneralStreamers)+len(res.InputedUsers); i++ {
		resp := <-channel
		imageMap[resp.Name] = resp.Image
	}
	iCount := len(res.InputedUsers)
	count := 0
	for i := 240.0; i < 240+iHeight; i += inputUserCardHeight {
		for j := 20.0; j < Width; j += 220 {
			user := res.InputedUsers[count]
			utils.DrawInputedUsers(c, j, i, 200, inputUserCardHeight-20, user, imageMap)
			count++
			if iCount == count {
				break
			}
		}
		if iCount == count {
			break
		}
	}
	count = 0
	for i := 240.0 + iHeight; i < Height; i += cardHeight {
		for j := 20.0; j < Width; j += 220 {
			streamer := res.GeneralStreamers[count]
			utils.DrawFragment(c, j, i, 200, cardHeight-20, streamer.Streamer, userMap[streamer.OldestUser.Username], streamer.OldestUser.Date, imageMap)
			count++
			if followsCount == count {
				break
			}
		}
		if followsCount == count {
			break
		}
	}

	c.SavePNG("result.png")
	return c.Image(), nil
}
