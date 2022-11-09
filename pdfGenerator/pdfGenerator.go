package pdfgenerator

type PDFGeneratorInterface interface {
	Create(htmlFile string) (string, error)
}
