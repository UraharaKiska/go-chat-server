package chatUser

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	// "github.com/UraharaKiska/go-auth/internal/client/db"
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/go-chat-server/internal/repository"
)

const (
	tableName             = "chat_user"
	chatIdColumn		  = "chat_id"
	usernameColumn		 = "username"
)

type repo struct {
	db db.Client
}


func NewRepository(db db.Client) repository.ChatUserRepository {
	return &repo{db: db}
}

func (r *repo) AddUsers(ctx context.Context, chatId int64, users *[]string) (error) {
	log.Printf("CHAT USER REPOSITORY - ADD USERS")
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIdColumn, usernameColumn)
	
	for _, username := range *users {
		builderInsert = builderInsert.Values(chatId, username)
	}

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "chatUser_repository.AddUsers",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}


