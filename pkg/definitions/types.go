package definitions

import (
	"flag"
	"github.com/spf13/viper"
)

type Config struct {
	Viper *viper.Viper
	Flags []flag.Flag
}

type Spec struct {
	Input         []string
	Asserts       []string
	CustomImports []string
}

type Case struct {
	SourceRoot   string
	FunctionName string
	Spec         Spec
	Module       string
}

type Cases struct {
	SourceRoot string
	Cases      []Case
}
