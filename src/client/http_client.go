package client

import (
	"image"
	"net/http"

	pkg "github.com/modaniru/image-generator/pkg/proto"
	"github.com/modaniru/image-generator/src/draw"
)

type HttpClient struct {
	client *http.Client
}

type ImageResponse struct {
	Image image.Image
	Name  string
}

func NewHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{client: client}
}

func (h *HttpClient) GetImageFromURI(user *pkg.ResponseUser, channel chan *draw.TwitchUser) {
	resp, err := h.client.Get(user.ImageLink)
	if err != nil {
		channel <- nil
		return
	}
	image, _, err := image.Decode(resp.Body)
	if err != nil {
		channel <- nil
		return
	}
	channel <- &draw.TwitchUser{
		ProfileImage: image,
		Username:     user.DisplayName,
		StreamerType: user.BroadcasterType,
	}
}
