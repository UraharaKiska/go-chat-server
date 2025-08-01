package auth

import (
	"github.com/UraharaKiska/go-chat-server/internal/service"
	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
	descAccess "github.com/UraharaKiska/go-chat-server/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
	accessClient descAccess.AccessV1Client
}

func NewImplementation(
		chatService service.ChatService,
		accessClient descAccess.AccessV1Client,
	) *Implementation {
	return &Implementation{
		chatService: chatService,
		accessClient: accessClient,
	}
}
