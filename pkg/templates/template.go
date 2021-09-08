package templates

import (
	"embed"
	"fmt"
	"log"
	"text/template"

	"github.com/r0x0d/testgen/internal/global"
	"github.com/r0x0d/testgen/pkg/definitions"
	"github.com/r0x0d/testgen/pkg/templates/python"
)

type Generator map[string]func(cases definitions.Cases, temp *template.Template) error

//go:embed python/*.tmpl
var templateFiles embed.FS

func BuildTemplate(cases definitions.Cases) error {
	temp, err := template.New(global.Config.Viper.GetString(definitions.ConfigKeyLanguage)).Parse(templateString())
	if err != nil {
		return err
	}

	generator := buildGenerator()
	err = generator[global.Config.Viper.GetString(definitions.ConfigKeyLanguage)](cases, temp)
	if err != nil {
		return err
	}

	return nil
}

func templateString() string {
	out, err := templateFiles.ReadFile(
		fmt.Sprintf(
			"%s/%s%s",
			global.Config.Viper.GetString(definitions.ConfigKeyLanguage),
			global.Config.Viper.GetString(definitions.ConfigKeyLanguage),
			definitions.TemplateExtension),
	)
	if err != nil {
		log.Fatalf("fatal: %v", err)
	}

	return string(out)
}

func buildGenerator() Generator {
	return Generator{global.Config.Viper.GetString(definitions.ConfigKeyLanguage): python.Build}
}
