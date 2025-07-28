package repository

import (
	"context"

	"github.com/UraharaKiska/go-chat-server/internal/model"
)


type ChatRepository interface {
	Create(ctx context.Context, chatInfo *model.ChatInfo) (int64, error)
	Delete(ctx context.Context, id int64) (error)
}

type ChatUserRepository interface {
	AddUsers(ctx context.Context, chatId int64, users *[]string) (error)
}

type ChatMessageRepository interface {
	AddMessage(ctx context.Context, message *model.MessageInfo) (error)
}



