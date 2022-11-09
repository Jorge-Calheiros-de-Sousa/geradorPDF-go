package pdfgenerator

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wk struct {
	rootPath string
}

func NewWkhHtmlToPdf(rootPath string) PDFGeneratorInterface {
	return &wk{rootPath: rootPath}
}

func (w *wk) Create(htmlFile string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "Cant Open HTML File", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	err = pdfg.Create()

	if err != nil {
		return "Cant Create", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	err = pdfg.WriteFile(fileName)

	if err != nil {
		return "Cant Write File", err
	}

	return fileName, nil
}
