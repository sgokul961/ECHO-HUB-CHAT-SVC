package interfaceR

import (
	"time"

	"github.com/sgokul961/echo-hub-chat-svc/pkg/domain"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRepository interface {
	CreateChatRoom(user1, user2 int64) error
	GetAllChats(userId int64) ([]models.Chat, error)
	GetMessages(id primitive.ObjectID) ([]domain.Messages, error)
	UpdateLastMessageAndTime(chatId primitive.ObjectID, lastMessage string, time time.Time) error
	IsChatExist(user1, user2 int64) (bool, error)
	IsValidChatId(chatId primitive.ObjectID) (bool, error)
	SaveMessage(message domain.Messages) (primitive.ObjectID, error)
	ReadMessage(chatId primitive.ObjectID, senderId int64) (int64, error)
	FetchRecipient(chatId primitive.ObjectID, userId int64) (int64, error)
	DeleteChatsAndMessagesByUserID(userID uint) error
}
