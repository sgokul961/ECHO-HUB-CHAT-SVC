package response

import "github.com/sgokul961/echo-hub-chat-svc/pkg/utils/models"

type ChatResponse struct {
	Chat models.Chat
	User models.UserShortDetail
}
