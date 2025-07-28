package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/UraharaKiska/go-chat-server/internal/model"
	"github.com/UraharaKiska/go-chat-server/internal/repository"
	repositoryMocks "github.com/UraharaKiska/go-chat-server/internal/repository/mock"
	"github.com/UraharaKiska/go-chat-server/internal/service/chat"
	"github.com/UraharaKiska/platform-common/pkg/db"
	txMock "github.com/UraharaKiska/platform-common/pkg/db/mock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSendMessage(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type chatUserRepositoryMockFunc func(mc *minimock.Controller) repository.ChatUserRepository
	type chatMessageRepositoryFunc func(mc *minimock.Controller) repository.ChatMessageRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.MessageInfo
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		from = gofakeit.Name()
		text = gofakeit.BeerName()
		date, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")

		transactionErr = fmt.Errorf("transaction error")

		message = &model.MessageInfo{
			From: from,
			Text: text,
			Timestamp: date,
		}

	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		err             error
		chatRepositoryMock chatRepositoryMockFunc
		chatUserRepositoryMock chatUserRepositoryMockFunc
		chatMessageRepositoryMock chatMessageRepositoryFunc
		txManagerMock txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: message,
			},
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repository.ChatUserRepository {
				mock := repositoryMocks.NewChatUserRepositoryMock(mc)
				return mock
			},
			chatMessageRepositoryMock: func(mc *minimock.Controller) repository.ChatMessageRepository {
				mock := repositoryMocks.NewChatMessageRepositoryMock(mc)
				mock.AddMessageMock.Expect(ctx, message).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "transaction error case",
			args: args{
				ctx: ctx,
				req: message,
			},
			err:  transactionErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)

				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repository.ChatUserRepository {
				mock := repositoryMocks.NewChatUserRepositoryMock(mc)
				return mock
			},
			chatMessageRepositoryMock: func(mc *minimock.Controller) repository.ChatMessageRepository {
				mock := repositoryMocks.NewChatMessageRepositoryMock(mc)
				mock.AddMessageMock.Expect(ctx, message).Return(transactionErr)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) db.TxManager {
				mock := txMock.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f db.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			chatRepositoryMock := tt.chatRepositoryMock(mc)
			chatUserRepositoryMock := tt.chatUserRepositoryMock(mc)
			chatMessageRepositoryMock := tt.chatMessageRepositoryMock(mc)
			txManagerMock := tt.txManagerMock(mc)
			service := chat.NewService(
				chatRepositoryMock, 
				chatMessageRepositoryMock,
				chatUserRepositoryMock, txManagerMock)


			err := service.SendMessage(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}

}
