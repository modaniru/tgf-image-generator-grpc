package service

import (
	"context"
	"errors"
	"image"

	pkg "github.com/modaniru/image-generator/pkg/proto"
	"github.com/modaniru/image-generator/src/client"
	"github.com/modaniru/image-generator/src/draw"
)

type ImageGeneratorService struct {
	tgfClient     *client.TgfClient
	httpClient    *client.HttpClient
	drawerService *DrawerService
}

func NewService(tgfClient *client.TgfClient, httpClient *client.HttpClient, drawService *DrawerService) *ImageGeneratorService {
	return &ImageGeneratorService{
		drawerService: drawService,
		tgfClient:     tgfClient,
		httpClient:    httpClient,
	}
}

// TODO Refactor
func (im *ImageGeneratorService) GenerateImage(nicknames []string) (image.Image, error) {
	response, err := im.tgfClient.GetGeneralFollows(context.Background(), &pkg.GetTGFRequest{Usernames: nicknames})
	if err != nil {
		return nil, err
	}
	tu, gs, err := im.mapping(response)
	if err != nil {
		return nil, err
	}
	return im.drawerService.DrawImage(tu, gs)
}

// Маппинг результата, для скачивания картинок
func (im *ImageGeneratorService) mapping(response *pkg.GetTGFResponse) ([]*draw.TwitchUser, []*draw.Streamer, error) {
	channel := make(chan *draw.TwitchUser)
	inputedUsersMap := make(map[string]*draw.TwitchUser)
	for _, i := range response.InputedUsers {
		go im.httpClient.GetImageFromURI(i, channel)
	}
	for range response.InputedUsers {
		u := <-channel
		if u == nil {
			return nil, nil, errors.New("user load immage error")
		}
		inputedUsersMap[u.Username] = u
	}
	streamersMap := make(map[string]*draw.Streamer)
	for _, s := range response.GeneralStreamers {
		streamersMap[s.Streamer.DisplayName] = &draw.Streamer{
			Oldest: inputedUsersMap[s.OldestUser.Username],
			Date:   s.OldestUser.Date,
		}
		go im.httpClient.GetImageFromURI(s.Streamer, channel)
	}
	for range response.GeneralStreamers {
		u := <-channel
		if u == nil {
			return nil, nil, errors.New("user load immage error")
		}
		streamersMap[u.Username].TwitchUser = u
	}
	inputedUsers := make([]*draw.TwitchUser, 0, len(inputedUsersMap))
	for _, v := range inputedUsersMap {
		inputedUsers = append(inputedUsers, v)
	}
	generalStreamers := make([]*draw.Streamer, 0, len(streamersMap))
	for _, v := range streamersMap {
		generalStreamers = append(generalStreamers, v)
	}
	return inputedUsers, generalStreamers, nil
}
