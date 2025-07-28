package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/UraharaKiska/go-chat-server/internal/api/chat"
	"github.com/UraharaKiska/go-chat-server/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/UraharaKiska/go-chat-server/internal/service"
	serviceMocks "github.com/UraharaKiska/go-chat-server/internal/service/mock"
	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatServiceMockFunc func(mc *minimock.Controller) service.ChatService
	
	type args struct {
		ctx context.Context
		req *desc.SendMessageRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		from = gofakeit.Name()
		text = gofakeit.BeerName()
		invalidDateStr = "ultra giga mega penis"
		dateStr = gofakeit.Date().Format(time.RFC3339)
		date, _ = time.Parse(time.RFC3339, dateStr)

		serviceErr = fmt.Errorf("service error")
		_, parseErr = time.Parse(time.RFC3339, invalidDateStr)

		req = &desc.SendMessageRequest{
			Message: &desc.MessageInfo{
				From: from,
				Text: text,
				Datetime: dateStr,
			},
		}
		reqInvalid = &desc.SendMessageRequest{
			Message: &desc.MessageInfo{
				From: from,
				Text: text,
				Datetime: invalidDateStr,
			},
		}
		
		message = &model.MessageInfo{
			From: from,
			Text: text,
			Timestamp: date,
		}

		res = &emptypb.Empty{}
	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
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
				mock.SendMessageMock.Expect(ctx, message).Return(nil)
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
				mock.SendMessageMock.Expect(ctx, message).Return(serviceErr)
				return mock
			},
		},
		{
			name: "error case invalid date format",
			args: args{
				ctx: ctx,
				req: reqInvalid,
			},
			want: nil,
			err:  parseErr,
			chatServiceMock: func(mc *minimock.Controller) service.ChatService {
				mock := serviceMocks.NewChatServiceMock(mc)
				// mock.SendMessageMock.Return(nil)

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

			sendMesResp, err := api.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, sendMesResp)
		})
	}

}
