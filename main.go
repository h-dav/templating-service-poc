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

	f, _ := os.Create("example_result-push.txt")

	tmpl.Execute(f, templateValues)
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

	f, _ := os.Create("example_result-email.html")

	tmpl.Execute(f, templateValues)
}
