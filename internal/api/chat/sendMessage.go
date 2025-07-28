package auth

import (
	"context"
	"log"

	// "github.com/UraharaKiska/go-auth/internal/converter"
	"github.com/UraharaKiska/go-chat-server/internal/converter"
	"github.com/UraharaKiska/go-chat-server/internal/utils"
	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("API - GET")
	_, err := utils.ParseDateTime(req.GetMessage().GetDatetime())
	if err != nil {
		return nil, err
	}
	err = i.chatService.SendMessage(ctx, converter.ToMessageInfoFromDesc(req.GetMessage()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil

}

