package server

import (
	"bytes"
	"context"
	"image/png"

	pkg "github.com/modaniru/image-generator/pkg/proto"
	"github.com/modaniru/image-generator/src/service"
)

type ImageServiceServer struct {
	pkg.ImageServiceServer
	service *service.ImageGeneratorService
}

func NewImageServiceServer(service *service.ImageGeneratorService) *ImageServiceServer {
	return &ImageServiceServer{service: service}
}

func (i *ImageServiceServer) GetGeneralFollowsImage(c context.Context, request *pkg.GeneralFollowsImageRequest) (*pkg.GeneralFollowsImageResponse, error) {
	image, err := i.service.GenerateImage(request.GetNicknames())
	if err != nil {
		return nil, err
	}
	// Image to []byte
	buffer := bytes.Buffer{}
	png.Encode(&buffer, image)
	return &pkg.GeneralFollowsImageResponse{
		Image: buffer.Bytes(),
	}, nil
}
