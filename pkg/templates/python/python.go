package python

import (
	"fmt"
	"github.com/r0x0d/testgen/internal/global"
	"github.com/r0x0d/testgen/pkg/definitions"
	"os"
	"text/template"
)

func BuildTemplate(cases definitions.Cases) error {
	temp, err := template.New(global.Config.Viper.GetString("language")).Parse(templateString())
	if err != nil {
		return err
	}

	file, _ := os.Create(fmt.Sprintf("%s/test_%s.py", global.Config.Flags[0].DefValue, cases.Cases[0].FunctionName))
	return executeTemplate(temp, cases, file)
}

func executeTemplate(temp *template.Template, cases definitions.Cases, file *os.File) error {
	if len(cases.Cases) == 0 {
		return nil
	}
	cases.Cases[0].SourceRoot = cases.SourceRoot
	err := temp.Execute(file, cases.Cases[0])
	if err != nil {
		return err
	}

	cases.Cases = cases.Cases[1:]
	return executeTemplate(temp, cases, file)
}

func templateString() string {
	return `
from {{.SourceRoot}}.{{.Module}} import {{.FunctionName}}

{{range $import := .Spec.CustomImports}}
{{$import}}
{{end}}

def test_{{.FunctionName}}():
	# Execution
	result = {{.FunctionName}}(
		{{range $value := .Spec.Input}}
		{{$value}},
		{{end}}
	)

	# Assertation
	{{range $value := .Spec.Asserts}}
	assert {{$value}}
	{{end}}
`
}
