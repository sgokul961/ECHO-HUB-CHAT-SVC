package clientinterface

import "github.com/sgokul961/echo-hub-chat-svc/pkg/utils/models"

type AuthServiceClient interface {
	FetchDetails(id int64) (models.UserShortDetail, error)
}
