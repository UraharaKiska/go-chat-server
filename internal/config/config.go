package config

import (
	"flag"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

type GRPCConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type HTTPConfig interface {
	Address() string
}

type SWAGGERConfig interface {
	Address() string
}

type TlsConfig interface {
	ServiceKeyFilePath() string
	ServicePemFilePath() string
}

func ParseConfig() string {
	var configPath string
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
	flag.Parse()
	return configPath
}
