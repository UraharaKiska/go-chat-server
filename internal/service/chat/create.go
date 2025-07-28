package chat

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	log.Printf("SERVICE - CREATE")
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepository.Create(ctx, &chat.Info)
		if errTx != nil {
			return errTx
		}
		errTx = s.chatUserRepository.AddUsers(ctx, id, &chat.Users)
		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return 0, err
	}
	return id, nil
}
