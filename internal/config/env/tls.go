package config

import (
	"github.com/UraharaKiska/go-chat-server/internal/config"
)


type tlsConfig struct {
	serviceKeyFilePath string
	servicePemFilePath string
}

func NewTLSConfig() (config.TlsConfig, error) {
	return &tlsConfig{
		serviceKeyFilePath: "service.key",
		servicePemFilePath: "service.pem",
	}, nil
}

func (cfg *tlsConfig) ServiceKeyFilePath() string{
	return cfg.serviceKeyFilePath
}

func (cfg *tlsConfig) ServicePemFilePath() string{
	return cfg.servicePemFilePath

}
