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
	csvFile := flag.String("c", "list.csv", "The csv file to parse")                   //O(1)-T O(1)-S
	templateName := flag.String("t", "report.tpl", "The template to use")              //O(1)-T O(1)-S
	fileNameColumn := flag.Int("i", 0, "Column of csv to use as output file basename") //O(1)-T O(1)-S
	fileExt := flag.String("s", "eml", "Output file suffix")                           //O(1)-T O(1)-S
	flag.Parse()                                                                       //O(1)-T O(n)-S

	template, err := template.ParseFiles(*templateName) //O(1)-T O(n)-S n - size of the template
	if err != nil {
		panic(err)
	} //O(1)-T O(1)-S

	file, err := os.Open(*csvFile) //O(1)-T O(n)-S
	if err != nil {
		panic(err)
	} //O(1)-T O(1)-S
	defer file.Close() //O(1)-T O(1)-S note space is released by the garbage collector
	e
	reader := csv.NewReader(file) //O(1)-T O(n)-S
	reader.Comma = ','
	for {
		row, err := reader.Read() //O(1)-T O(n)-S number of column
		if err == io.EOF {
			break //O(1)-T O(1)-S
		} else if err != nil {
			panic(err)
		} //O(1)-T O(1)-S
		f, err := os.Create("./output/" + row[*fileNameColumn] + "." + *fileExt) //O(1)-T O(1)-S

		defer f.Close()                          //O(1)-T O(1)-S note space is released by the garbage collector
		template.Execute(f, TWrap{Fields: &row}) //O(1)-T O(m)-S m - number of template slots
	} //O(n)-T O(n+m)-S
	sendEmail()
}
