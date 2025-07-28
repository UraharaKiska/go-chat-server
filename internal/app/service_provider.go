package app

import (
	"context"
	"log"
	"github.com/UraharaKiska/go-chat-server/internal/api/chat"
	"github.com/UraharaKiska/platform-common/pkg/db"
	"github.com/UraharaKiska/platform-common/pkg/db/pg"
	"github.com/UraharaKiska/platform-common/pkg/db/transaction"
	"github.com/UraharaKiska/platform-common/pkg/closer"
	"github.com/UraharaKiska/go-chat-server/internal/config"
	env "github.com/UraharaKiska/go-chat-server/internal/config/env"
	"github.com/UraharaKiska/go-chat-server/internal/repository"
	chatRepository "github.com/UraharaKiska/go-chat-server/internal/repository/chat"
	chatMessageRepository "github.com/UraharaKiska/go-chat-server/internal/repository/chatMessage"
	chatUserRepository "github.com/UraharaKiska/go-chat-server/internal/repository/chatUser"

	"github.com/UraharaKiska/go-chat-server/internal/service"
	chatService "github.com/UraharaKiska/go-chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager		  	db.TxManager
	chatRepository repository.ChatRepository
	chatMessageRepository repository.ChatMessageRepository
	chatUserRepository repository.ChatUserRepository


	chatService service.ChatService

	authImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to init app: #{err.Error()}")
		}
		s.pgConfig = cfg
	}
	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to init app: #{err.Error()}")
		}
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to init app: #{err.Error()}")
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to init app: #{err.Error()}")
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}
	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactorManager(s.DBClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatMessageRepository(ctx context.Context) repository.ChatMessageRepository {
	if s.chatMessageRepository == nil {
		s.chatMessageRepository = chatMessageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatMessageRepository
}

func (s *serviceProvider) ChatUserRepository(ctx context.Context) repository.ChatUserRepository {
	if s.chatUserRepository == nil {
		s.chatUserRepository = chatUserRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatUserRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.ChatMessageRepository(ctx),
			s.ChatUserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.ChatService(ctx))
	}

	return s.authImpl
}
