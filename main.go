package main

import (
	"encoding/csv"
	"flag"
	"io"
	"os"
	"text/template"
)

type TWrap struct{ Fields *[]string }

func main() {
	csvFile := flag.String("c", "list.csv", "The csv file to parse")
	templateName := flag.String("t", "report.tpl", "The template to use")
	fileNameColumn := flag.Int("i", 0, "Column of csv to use as output file basename")
	fileExt := flag.String("s", "eml", "Output file suffix")
	flag.Parse()

	// loading the template
	template, err := template.ParseFiles(*templateName)
	if err != nil {
		panic(err)
	}

	// loading the csv file for data
	file, err := os.Open(*csvFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// parsing the csv, one record at a time
	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		// get next row... exit on EOF
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		f, err := os.Create("./output/" + row[*fileNameColumn] + "." + *fileExt)

		defer f.Close()
		template.Execute(f, TWrap{Fields: &row})
	}
	sendEmail()
}
