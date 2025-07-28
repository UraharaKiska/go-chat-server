package auth

import (
	"context"
	"log"

	// "github.com/UraharaKiska/go-auth/internal/converter"
	"github.com/UraharaKiska/go-chat-server/internal/converter"
	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("API - GET")
	id, err := i.chatService.Create(ctx, converter.ToChatFromDesc(req))
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{
		Id: id,
	}, nil

}
