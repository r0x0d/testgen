package global

import (
	"flag"
	"github.com/r0x0d/testgen/pkg/definitions"
	"github.com/spf13/viper"
)

var (
	Config = definitions.Config{
		Viper: viper.New(),
		Flags: []flag.Flag{},
	}
)
