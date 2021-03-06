package main // import "github.com/hectorj/go-resultgen"

import (
	"bytes"
	"flag"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type templateData struct {
	PkgName             string
	GeneratedStructName string
	TypeName            string
	StrictMode          bool
	BuildTags           string
}

var tpl = template.Must(template.New("output").Parse(tplSrc))

func main() {
	var data templateData

	if len(os.Args) < 2 {
		log.Fatalln("Expected type name as first argument")
	}
	data.TypeName = os.Args[1]
	var outputPath string

	flagSet := flag.NewFlagSet("options", flag.ExitOnError)
	flagSet.BoolVar(&data.StrictMode, "strict", false, "")
	flagSet.StringVar(&data.BuildTags, "tags", "", "")
	flagSet.StringVar(&outputPath, "output", strings.ToLower(data.TypeName)+"_result.go", "")
	flagSet.Parse(os.Args[2:])

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	pkg, err := build.Default.ImportDir(wd, 0)
	if err != nil {
		log.Fatalln(err)
	}

	data.PkgName = pkg.Name
	data.GeneratedStructName = upperCaseFirst(data.TypeName + "Result")

	outputBuffer := bytes.NewBuffer(nil)
	err = tpl.Execute(outputBuffer, data)
	if err != nil {
		log.Fatalln(err)
	}

	if !filepath.IsAbs(outputPath) {
		outputPath = filepath.Join(wd, outputPath)
	}
	err = ioutil.WriteFile(outputPath, outputBuffer.Bytes(), 0666)
	if err != nil {
		log.Fatalln(err)
	}
}

func upperCaseFirst(s string) string {
	r := []rune(s)
	r = append([]rune(strings.ToUpper(string(r[0]))), r[1:]...)
	return string(r)
}

const tplSrc = `{{if ne .BuildTags "" }}// +build {{ .BuildTags }}

{{end}}package {{ .PkgName }}

import (
	"errors"
)

// File generated by github.com/hectorj/go-resultgen.

// {{ .GeneratedStructName }} is a result type for {{ .TypeName }}.
// See https://en.wikipedia.org/wiki/Result_type .
type {{ .GeneratedStructName }} struct {
	value {{ .TypeName }}
	err error{{if .StrictMode}}
	errWasChecked *bool{{end}}
}

func NewValid{{ .GeneratedStructName }}(value {{ .TypeName }}) {{ .GeneratedStructName }} {
	return {{ .GeneratedStructName }}{
		value: value,
		err: nil,{{if .StrictMode}}
		errWasChecked: new(bool),{{end}}
	}
}

func NewFailed{{ .GeneratedStructName }}(err error) {{ .GeneratedStructName }} {
	if err == nil {
		panic(errors.New("cannot create failed result from nil error"))
	}
	return {{ .GeneratedStructName }}{
		err: err,{{if .StrictMode}}
		errWasChecked: new(bool),{{end}}
	}
}

func (r {{ .GeneratedStructName }}) is{{ .GeneratedStructName }}() {}

func (r {{ .GeneratedStructName }}) GetError() error {{"{"}}{{if .StrictMode}}
	(*r.errWasChecked) = true{{end}}
	return r.err
}

func (r {{ .GeneratedStructName }}) Get{{ .TypeName }}() {{ .TypeName }} {{"{"}}{{if .StrictMode}}
	if !(*r.errWasChecked) {
		panic(errors.New("unsafe behavior: error was not checked before trying to get the value"))
	}{{end}}
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}
`
