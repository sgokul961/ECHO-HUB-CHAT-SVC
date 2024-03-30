package interfaceU

import (
	"github.com/sgokul961/echo-hub-chat-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/utils/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatUseCase interface {
	GetAllChats(userId int64) ([]response.ChatResponse, error)
	GetMessages(chatId primitive.ObjectID) ([]domain.Messages, error)
	SaveMessage(chatId primitive.ObjectID, senderId int64, message string) (primitive.ObjectID, error)
	ReadMessage(userId int64, chatId primitive.ObjectID) (int64, error)
	FetchRecipient(chatId primitive.ObjectID, userId int64) (int64, error)
	CreateChatRoom(user1, user2 int64) error
}
