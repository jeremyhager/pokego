package generate

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/jeremyhager/pokego/pkg/pokedex"
)

func RenderTemplate(inputFile, outputFile, id string) error {
	filePath := os.Stdout

	info, err := pokedex.Init(id)
	if err != nil {
		return err
	}

	tmplFile, err := template.ParseFiles(inputFile)
	if err != nil {
		return err
	}

	if outputFile == "" {
		filePath = os.Stdout
	} else {
		filePath, err = createPath(outputFile)
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
