package chat

import (
	"context"
	"log"

	"github.com/UraharaKiska/go-chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, messageInfo *model.MessageInfo) (error) {
	log.Printf("SERVICE - SendMessage")
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.chatMessageRepository.AddMessage(ctx, messageInfo)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
