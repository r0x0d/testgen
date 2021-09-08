package build

import (
	"flag"
	"fmt"

	"github.com/r0x0d/testgen/internal/config"
	"github.com/r0x0d/testgen/internal/global"
	"github.com/r0x0d/testgen/pkg/definitions"
)

func Prepare(flags []flag.Flag) error {
	err := config.Read(flags)
	if err != nil {
		return err
	}
	return nil
}

func Build() definitions.Cases {
	caseConfigList := global.Config.Viper.GetStringMapString("cases")
	var cases []definitions.Case
	for key, _ := range caseConfigList {
		cases = append(cases, definitions.Case{
			FunctionName: global.Config.Viper.GetString(fmt.Sprintf("cases.%s.function_name", key)),
			Module:       global.Config.Viper.GetString(fmt.Sprintf("cases.%s.module", key)),
			Spec: definitions.Spec{
				Input:   global.Config.Viper.GetStringSlice(fmt.Sprintf("cases.%s.spec.input", key)),
				Asserts: global.Config.Viper.GetStringSlice(fmt.Sprintf("cases.%s.spec.asserts", key)),
			},
		})
	}

	return definitions.Cases{SourceRoot: global.Config.Viper.GetString("source_root"), Cases: cases}
}
