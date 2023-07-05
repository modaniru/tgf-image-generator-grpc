package client

import (
	pkg "github.com/modaniru/image-generator/pkg/proto"
	"google.golang.org/grpc"
)

type TgfClient struct {
	pkg.TwitchGeneralFollowsClient
}

func NewTgfClient(conn *grpc.ClientConn) *TgfClient {
	return &TgfClient{pkg.NewTwitchGeneralFollowsClient(conn)}
}
