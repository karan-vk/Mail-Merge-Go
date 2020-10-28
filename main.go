package main

import ("encoding/csv"; "flag"; "io"; "os"; "text/template")

/// Wrap an array of strings as a struct, so we can pass it to a template
type TWrap struct { Fields *[]string }

/// Parse a CSV so that each line becomes an array of strings, and then use
/// the array of strings with a template to generate one file per csv line
func main() {
	// parse command line options
	csvFile := flag.String("c", "list.csv", "The csv file to parse")
	templateName := flag.String("t", "report.template", "The template to use")
	fileNameColumn := flag.Int("i", 0, "Column of csv to use as output file basename")
	fileExt := flag.String("s", "eml", "Output file suffix")
	flag.Parse()

	// loading the template
	template, err := template.ParseFiles(*templateName)
	if err != nil { panic(err) }

	// loading the csv file for data
	file, err := os.Open(*csvFile)
	if err != nil { panic(err) }
	defer file.Close()

	// parsing the csv, one record at a time
	reader := csv.NewReader(file)
	reader.Comma = ','
	for {
		// get next row... exit on EOF
		row, err := reader.Read()
		if err == io.EOF { break } else if err != nil { panic(err) }

		// create output file for row
		f, err := os.Create("./output/" + row[*fileNameColumn] + "." + *fileExt)
		//os.OpenFile(f)
		//if err != nil { panic(err) }
		defer f.Close()
		// apply template to row, dump to file
		template.Execute(f, TWrap{Fields: &row})
	}
	sendEmail()
}