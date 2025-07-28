package chat

import (
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-chat-server/internal/repository"
	"github.com/UraharaKiska/go-chat-server/internal/service"

)

type serv struct {
	chatRepository repository.ChatRepository
	chatMessageRepository repository.ChatMessageRepository
	chatUserRepository repository.ChatUserRepository
	txManager db.TxManager
}

func NewService(
	chatRepository repository.ChatRepository,
	chatMessageRepository repository.ChatMessageRepository,
	chatUserRepository repository.ChatUserRepository,
	txManager db.TxManager,
) service.ChatService {
	return &serv{
		chatRepository: chatRepository,
		chatMessageRepository: chatMessageRepository,
		chatUserRepository: chatUserRepository,
		txManager: txManager,
	}
}
