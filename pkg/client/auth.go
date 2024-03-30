package client

import (
	"context"
	"errors"
	"fmt"

	clientinterface "github.com/sgokul961/echo-hub-chat-svc/pkg/client/clientInterface"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/config"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/pb"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/utils/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient(c config.Config) clientinterface.AuthServiceClient {

	//if call fails chech for error here grpc communication

	cc, err := grpc.Dial(c.AuthHubUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("coudnt connect:", err)
	}
	return &AuthServiceClient{Client: pb.NewAuthServiceClient(cc)}
}
func (a *AuthServiceClient) FetchDetails(id int64) (models.UserShortDetail, error) {
	req := &pb.FetchShortDetailsRequest{
		Id: id,
	}

	res, err := a.Client.FetchShortDetails(context.Background(), req)
	if err != nil {
		return models.UserShortDetail{}, errors.New("error in getting auth service method")
	}
	response := models.UserShortDetail{
		Id:             id,
		Username:       res.Name,
		ProfilePicture: res.Image,
	}

	return response, nil
}
