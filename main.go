package main

import (
	"fmt"
	htmlparser "geradorPDF/htmlParser"
	pdfgenerator "geradorPDF/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {
	h := htmlparser.New("tmp")
	wk := pdfgenerator.NewWkhHtmlToPdf("tmp")
	dataHTML := Data{
		Name: "Jorge Calheiros de Sousa",
	}

	htmlGenerated, err := h.Create("templates/exemple.html", dataHTML)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("HTML gerado: ", htmlGenerated)

	filePDFName, err := wk.Create(htmlGenerated)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PDF gerado: ", filePDFName)
}
