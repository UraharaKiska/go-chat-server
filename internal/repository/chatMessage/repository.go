package chatMessage

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	// "github.com/UraharaKiska/go-auth/internal/client/db"
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-chat-server/internal/model"
	"github.com/UraharaKiska/go-chat-server/internal/repository"
)

const (
	tableName             = "chat_message"
	fromUserColumn		= "from_user"
	messageColumn		 = "message"
	createdAtColumn       = "created_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatMessageRepository {
	return &repo{db: db}
}

func (r *repo) AddMessage(ctx context.Context, message *model.MessageInfo) (error) {
	log.Printf("REPOSITORY - CREATE")
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(fromUserColumn, messageColumn, createdAtColumn ).
		Values(message.From, message.Text, message.Timestamp)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chatMessage_repository.AddMessage",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}


