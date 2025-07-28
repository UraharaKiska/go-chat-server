package chat

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
	tableName             = "chat"
	idColumn              = "id"
	nameColumn			  = "name"
	createdAtColumn       = "created_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	log.Printf("REPOSITORY - CREATE")
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn).
		Values(info.Name).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var chatId int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&chatId)
	if err != nil {
		return 0, err
	}

	return chatId, nil
}


func (r *repo) Delete(ctx context.Context, id int64) (error) {
	log.Printf("REPOSITORY - DELETE")
	builderDelete := sq.Delete(tableName).
	PlaceholderFormat(sq.Dollar).
	Where(sq.Eq{idColumn: id})
	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		Name:     "chat_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}


