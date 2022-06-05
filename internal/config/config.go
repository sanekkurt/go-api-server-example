package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/jessevdk/go-flags"
)

var (
	opts struct {
		ConfigPath string `long:"config" short:"c" env:"CONFIG_PATH" description:"Path to config.yaml file" required:"true"`
	}
	ErrHelpShown = errors.New("help message shown")

	globalCfg *AppConfig

)

func Parse(ctx context.Context, args []string) (*AppConfig, error) {
	_, err := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash).ParseArgs(args[1:])
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				return nil, ErrHelpShown
			}

			return nil, fmt.Errorf("cannot parse arguments: %w", flagsErr)
		}

		return nil, fmt.Errorf("cannot parse arguments: %w", err)
	}

	cfg, err := ParseConfig(ctx, opts.ConfigPath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return cfg, nil
}

func GetConfig() AppConfig {
	if globalCfg == nil {
		return AppConfig{}
	}

	return *globalCfg
}
