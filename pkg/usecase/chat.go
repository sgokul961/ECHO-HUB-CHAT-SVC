package usecase

import (
	"errors"
	"fmt"
	"time"

	clientinterface "github.com/sgokul961/echo-hub-chat-svc/pkg/client/clientInterface"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/domain"
	interfaceR "github.com/sgokul961/echo-hub-chat-svc/pkg/repository/interface"
	interfaceU "github.com/sgokul961/echo-hub-chat-svc/pkg/usecase/interface"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/utils/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatUseCase struct {
	Repo       interfaceR.ChatRepository
	AuthClient clientinterface.AuthServiceClient
}

func NewChatUseCase(repo interfaceR.ChatRepository, authClient clientinterface.AuthServiceClient) interfaceU.ChatUseCase {
	return &ChatUseCase{Repo: repo,
		AuthClient: authClient,
	}
}

func (c *ChatUseCase) GetAllChats(userId int64) ([]response.ChatResponse, error) {
	chats, err := c.Repo.GetAllChats(userId)
	if err != nil {
		return nil, err
	}
	chatResponses := make([]response.ChatResponse, 0)
	for _, chat := range chats {
		userIdInt64 := int64(chat.Users[0])

		user, err := c.AuthClient.FetchDetails(userIdInt64)
		if err != nil {
			return nil, err
		}
		chatResponse := response.ChatResponse{
			Chat: chat,
			User: user,
		}
		chatResponses = append(chatResponses, chatResponse)
	}
	return chatResponses, nil
}

func (c *ChatUseCase) GetMessages(chatId primitive.ObjectID) ([]domain.Messages, error) {
	isValid, err := c.Repo.IsValidChatId(chatId)
	if err != nil {
		return nil, err
	}
	if !isValid {
		return nil, errors.New("chatId is not existing")
	}
	messages, err := c.Repo.GetMessages(chatId)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (c *ChatUseCase) SaveMessage(chatId primitive.ObjectID, senderId int64, message string) (primitive.ObjectID, error) {
	isValid, err := c.Repo.IsValidChatId(chatId)
	fmt.Println("is valid ", isValid)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	if !isValid {
		return primitive.ObjectID{}, errors.New("chatId is not existing")
	}
	messages := domain.Messages{
		SenderID:       senderId,
		ChatID:         chatId,
		MessageContent: message,
		Timestamp:      time.Now(),
	}
	res, err := c.Repo.SaveMessage(messages)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	err = c.Repo.UpdateLastMessageAndTime(chatId, message, messages.Timestamp)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res, nil
}

// func (c *ChatUseCase) ReadMessage(userId uint, chatId primitive.ObjectID) (int64, error) {

// 	isValid, err := c.Repo.IsValidChatId(chatId)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if !isValid {
// 		return 0, errors.New("chatId is not existing")
// 	}
// 	senderId, err := c.Repo.FetchRecipient(chatId, userId)

// 	if err != nil {
// 		return 0, err
// 	}

// 	res, err := c.Repo.ReadMessage(chatId, senderId)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return res, nil
// }

func (c *ChatUseCase) FetchRecipient(chatId primitive.ObjectID, userId int64) (int64, error) { // <--------------------changed
	isValid, err := c.Repo.IsValidChatId(chatId)
	if err != nil {
		return 0, err
	}
	if !isValid {
		return 0, errors.New("chatId is not existing")
	}
	res, err := c.Repo.FetchRecipient(chatId, userId)
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (c *ChatUseCase) CreateChatRoom(user1, user2 int64) error {
	create := c.Repo.CreateChatRoom(user1, user2)
	if create != nil {
		return create
	}
	return nil
}
