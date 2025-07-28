package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/UraharaKiska/go-chat-server/internal/api/chat"
	"github.com/UraharaKiska/go-chat-server/internal/model"

	"github.com/UraharaKiska/go-chat-server/internal/service"
	serviceMocks "github.com/UraharaKiska/go-chat-server/internal/service/mock"
	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	
	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = gofakeit.Int64()
		name             = gofakeit.Name()
		users = []string{ name, name, name}

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			ChatInfo: &desc.ChatInfo{
				Name:            name,
			},
			Usernames: users,
		}
		
		chat = &model.Chat{
			Info: model.ChatInfo{
				Name: name,
			},
			Users: users,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		chatServiceMock chatServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(id, nil)
				return mock
			},
		},
		{
			name: "success error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				mock.CreateMock.Expect(ctx, chat).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.chatServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}

}
