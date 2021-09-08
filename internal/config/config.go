package config

import (
	"flag"

	"github.com/r0x0d/testgen/internal/global"
)

func Read(flags []flag.Flag) error {
	global.Config.Flags = flags
	global.Config.Viper.SetConfigName("testgen")
	global.Config.Viper.SetConfigType("yaml")
	global.Config.Viper.AddConfigPath(".")
	err := global.Config.Viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
