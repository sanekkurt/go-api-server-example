package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go-api-server-example/internal/config/configstructs"
	"go-api-server-example/internal/logging"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	ConfigFilePath string

	Server configstructs.Server `yaml:"server"`
}

func ParseConfig(ctx context.Context, cfgPath string) (*AppConfig, error) {
	log := logging.GetLogger()

	configFullPath, err := filepath.Abs(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve absolute path to %s: %w", configFullPath, err)
	}

	f, err := os.Open(configFullPath) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("cannot open config '%s' for reading: %w", configFullPath, err)
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Warnf("cannot close config '%s': %s", cfgPath, err)
		}
	}()

	dec := yaml.NewDecoder(f)
	dec.SetStrict(true)

	cfg := &AppConfig{}

	err = dec.Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config from '%s': %w", cfgPath, err)
	}

	cfg.ConfigFilePath = filepath.Dir(configFullPath)

	return cfg, nil
}
