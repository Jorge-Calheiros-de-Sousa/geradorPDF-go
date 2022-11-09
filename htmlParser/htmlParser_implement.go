package htmlparser

import (
	"os"
	"text/template"

	"github.com/google/uuid"
)

type HTMLScrut struct {
	rootPath string
}

func New(rootPath string) HTMLParseInterface {
	return &HTMLScrut{rootPath: rootPath}
}

func (a *HTMLScrut) Create(templateName string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)

	if err != nil {
		return "Invalid!", err
	}

	fileName := a.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)

	if err != nil {
		return "Invalid!", err
	}

	err = templateGenerator.Execute(fileWriter, data)
	if err != nil {
		return "Invalid!", err
	}

	return fileName, nil
}
