package main

import (
	"context"
	"log"
	"github.com/UraharaKiska/go-chat-server/internal/app"
	// "github.com/brianvoe/gofakeit"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
	// "google.golang.org/protobuf/types/known/timestamppb"
	// "google.golang.org/protobuf/types/known/emptypb"
	// desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
)

// const grpcPort = 50051

// type server struct {
// 	desc.UnimplementedChatV1Server
// }

// func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
// 	usernames := req.GetUsernames()
// 	for id, user := range usernames {
// 		log.Printf("User %d:, %v", id, user)
// 	}
// 	return &desc.CreateResponse{
// 		Id: int64(gofakeit.Uint8()),
// 	}, nil
// }

// func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
// 	log.Printf("Chat id: %d", req.GetId())

// 	return &emptypb.Empty{}, nil
// }

// func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
// 	log.Printf("From: %+v", req.GetFrom())
// 	log.Printf("Text: %+v", req.GetText())
// 	log.Printf("Timestamp: %+v", req.GetTimestamp())
// 	return &emptypb.Empty{}, nil
// }

func main() {
		ctx := context.Background()
	a, err := app.NewApp(ctx)
	// log.Printf("App :%v", a)
	if err != nil {
		log.Fatalf("failed to init app%v: ", err)
	}
	err = a.Run()
	if err != nil {
		log.Fatalf("failed to init app%v: ", err)
	}
}