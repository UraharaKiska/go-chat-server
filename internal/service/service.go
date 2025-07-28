package service

import (
	"context"

	"github.com/UraharaKiska/go-chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	SendMessage(ctx context.Context, messageInfo *model.MessageInfo) (error)
	Delete(ctx context.Context, id int64) (error)
}
