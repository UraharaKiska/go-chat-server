package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/UraharaKiska/platform-common/pkg/db"
	txMock "github.com/UraharaKiska/platform-common/pkg/db/mock"
	"github.com/UraharaKiska/go-chat-server/internal/model"
	"github.com/UraharaKiska/go-chat-server/internal/repository"
	repositoryMocks "github.com/UraharaKiska/go-chat-server/internal/repository/mock"
	"github.com/UraharaKiska/go-chat-server/internal/service/chat"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type chatRepositoryMockFunc func(mc *minimock.Controller) repository.ChatRepository
	type chatUserRepositoryMockFunc func(mc *minimock.Controller) repository.ChatUserRepository
	type chatMessageRepositoryFunc func(mc *minimock.Controller) repository.ChatMessageRepository
	type txManagerMockFunc func(mc *minimock.Controller) db.TxManager

	type args struct {
		ctx context.Context
		req *model.Chat
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id               = gofakeit.Int64()
		name             = gofakeit.Name()
		users = []string{ name, name, name}

		transactionErr = fmt.Errorf("transaction error")

		req = &model.Chat{
			Info: model.ChatInfo{
				Name: name,
			},
			Users: users,
		}

	)

	// defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            int64
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
				req: req,
			},
			want: id,
			err:  nil,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &req.Info).Return(id, nil)
				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repository.ChatUserRepository {
				mock := repositoryMocks.NewChatUserRepositoryMock(mc)
				mock.AddUsersMock.Expect(ctx, id, &users).Return(nil)
				return mock
			},
			chatMessageRepositoryMock: func(mc *minimock.Controller) repository.ChatMessageRepository {
				mock := repositoryMocks.NewChatMessageRepositoryMock(mc)
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
				req: req,
			},
			want: 0,
			err:  transactionErr,
			chatRepositoryMock: func(mc *minimock.Controller) repository.ChatRepository {
				mock := repositoryMocks.NewChatRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &req.Info).Return(id, nil)
				return mock
			},
			chatUserRepositoryMock: func(mc *minimock.Controller) repository.ChatUserRepository {
				mock := repositoryMocks.NewChatUserRepositoryMock(mc)
				mock.AddUsersMock.Expect(ctx, id, &users).Return(transactionErr)
				return mock
			},
			chatMessageRepositoryMock: func(mc *minimock.Controller) repository.ChatMessageRepository {
				mock := repositoryMocks.NewChatMessageRepositoryMock(mc)
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


			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}

}
