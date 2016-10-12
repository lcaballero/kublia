package conf

import (
	"github.com/go-ini/ini"
)

type Keys struct {
	Environment string `ini:"environment"`
	Session     string `ini:"session"`
}

const (
	DevEnv          = "dev"
	DefaultKeysFile = "keys"
)

func LoadKeys(env string) (*Keys, error) {
	cfg, err := ini.Load("keys")
	if err != nil {
		return nil, err
	}

	keys := &Keys{
		Environment: env,
	}
	err = cfg.Section(env).MapTo(keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
