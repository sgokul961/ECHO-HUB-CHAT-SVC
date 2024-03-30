package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/sgokul961/echo-hub-chat-svc/pkg/pb"
	interfaceU "github.com/sgokul961/echo-hub-chat-svc/pkg/usecase/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHandler struct {
	UseCase interfaceU.ChatUseCase
	pb.UnimplementedChatServiceServer
}

func NewChatHandler(usecase interfaceU.ChatUseCase) *ChatHandler {
	return &ChatHandler{
		UseCase: usecase,
	}
}
func (c *ChatHandler) CreateChatRoom(ctx context.Context, createroom *pb.ChatRoomRequest) (*pb.ChatRoomResponse, error) {
	// Create chat room using the use case
	err := c.UseCase.CreateChatRoom(createroom.FollowingId, createroom.FOllowerId)

	// Check if there was an error
	if err != nil {
		// Return an error response
		return &pb.ChatRoomResponse{Error: err.Error()}, nil
	}

	// Return a success response
	return &pb.ChatRoomResponse{}, nil
}

func (c *ChatHandler) SaveMessage(ctx context.Context, p *pb.SaveMessageRequest) (*pb.SaveMessageResponse, error) {
	chatIdToHex, err := primitive.ObjectIDFromHex(p.ChatId)
	fmt.Println("chatidhex :", chatIdToHex)
	if err != nil {
		return &pb.SaveMessageResponse{}, err
	}

	res, err := c.UseCase.SaveMessage(chatIdToHex, p.UserId, p.Messages)

	// Check if there was an error
	if err != nil {
		// Return an error response
		return &pb.SaveMessageResponse{}, err
	}

	return &pb.SaveMessageResponse{
		Res: res.Hex(),
	}, nil

}
func (c *ChatHandler) FetchRecipient(ctx context.Context, p *pb.FetchRecipientRequest) (*pb.FetchRecipientResponse, error) {
	chatIdToHex, err := primitive.ObjectIDFromHex(p.ChatId)
	if err != nil {
		return &pb.FetchRecipientResponse{}, err
	}

	res, err := c.UseCase.FetchRecipient(chatIdToHex, p.UserId)

	if err != nil {
		return &pb.FetchRecipientResponse{}, err
	}

	return &pb.FetchRecipientResponse{
		Recipient: res,
	}, nil

}
func (c *ChatHandler) Getchats(ctx context.Context, p *pb.GetchatsRequest) (*pb.GetchatsResponse, error) {

	res, err := c.UseCase.GetAllChats(p.UserId)

	if err != nil {
		return &pb.GetchatsResponse{}, err
	}

	var chatResponses []*pb.ChatResponse

	for _, chat := range res {
		pbChat := &pb.ChatResponse{
			Chat: &pb.Chat{
				Users:           chat.Chat.Users,
				LastMessage:     chat.Chat.LastMessage,
				LastMessageTime: chat.Chat.LastMessageTime.String(),
			},
			User: &pb.UserShortDetail{
				Id:    chat.User.Id,
				Name:  chat.User.Username,
				Image: chat.User.ProfilePicture,
			},
		}
		chatResponses = append(chatResponses, pbChat)

	}
	return &pb.GetchatsResponse{
		ChatResponses: chatResponses,
	}, nil

}
func (c *ChatHandler) GetMessages(ctx context.Context, p *pb.GetMessagesRequest) (*pb.GetMessagesReponse, error) {

	chatid, err := primitive.ObjectIDFromHex(p.ChatId)

	if err != nil {
		return &pb.GetMessagesReponse{}, err
	}

	res, err := c.UseCase.GetMessages(chatid)

	if err != nil {
		return &pb.GetMessagesReponse{}, err
	}

	var getMessages []*pb.Message

	for _, msg := range res {
		pbMsg := &pb.Message{
			Id:              msg.ID.Hex(),
			SenderId:        msg.SenderID,
			ChatId:          msg.ChatID.Hex(),
			Seen:            msg.Seen,
			Image:           msg.Image,
			MessageContent:  msg.MessageContent,
			LastMessageTime: msg.Timestamp.Format(time.RFC3339), // Convert time.Time to string in RFC3339 format
		}
		getMessages = append(getMessages, pbMsg)
	}

	return &pb.GetMessagesReponse{
		Messages: getMessages,
	}, nil
}
