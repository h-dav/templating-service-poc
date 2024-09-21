package main

import (
	htmlTmpl "html/template"
	"os"
	textTmpl "text/template"
)

type ExampleValues struct {
	FirstName string
}

func main() {
	handleSimpleString()
	handleHtml()
}

func handleSimpleString() {
	templateFile := "example-push.txt"

	tmpl, _ := textTmpl.New(templateFile).ParseFiles(templateFile)

	templateValues := ExampleValues{
		FirstName: "Harry",
	}

	tmpl.Execute(os.Stdout, templateValues)
}

type ExampleHtmlValues struct {
	PageTitle string
	FirstName string
	Interest  struct {
		Gross string
		AER   string
	}
}

func handleHtml() {
	templateFile := "example-email.html"

	templateValues := ExampleHtmlValues{
		PageTitle: "Welcome",
		FirstName: "Harry",
	}

	templateValues.Interest.Gross = "2.35"
	templateValues.Interest.AER = "2.36"

	tmpl, _ := htmlTmpl.New(templateFile).ParseFiles(templateFile)

	tmpl.Execute(os.Stdout, templateValues)

}
