package cli

import (
	"github.com/jessevdk/go-flags"
	"github.com/lcaballero/kublai/conf"
	"os"
)

func ParseArgs(args ...string) *conf.Config {
	cfg := &conf.Config{}
	parser := flags.NewParser(cfg, flags.Default)
	_, err := parser.ParseArgs(args)
	if err != nil {
		os.Exit(1)
	}
	return cfg
}
