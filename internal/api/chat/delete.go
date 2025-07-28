package auth

import (
	"context"
	"log"

	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("API - DELETE")
	err := i.chatService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
