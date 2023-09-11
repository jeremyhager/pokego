package render

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/jeremyhager/pokego/internal/pokemoninfo"
)

type RenderArgs struct {
	ID         string
	InputFile  string
	OutputFile string
	Debug      bool
}

func (r *RenderArgs) RenderTemplate() error {
	filePath := os.Stdout

	info, err := pokemoninfo.Init(r.ID)
	if err != nil {
		return err
	}

	tmplFile, err := template.ParseFiles(r.InputFile)
	if err != nil {
		return err
	}

	if r.OutputFile == "" {
		filePath = os.Stdout
	} else {
		filePath, err = createPath(r.OutputFile)
		if err != nil {
			return err
		}
	}

	err = tmplFile.Execute(filePath, info)

	if err != nil {
		return err
	}
	return nil
}

func createPath(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return nil, err
	}
	return os.Create(path)
}
