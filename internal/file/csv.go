package file

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
) 

func ReadCSVFile(filename string) [][] string{
	
	fileData, err := ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	parsedFile := string(fileData)
	r := csv.NewReader(strings.NewReader(parsedFile))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func WriteCSVFile(filename string, data [][] string) {
	file, err := os.OpenFile(baseFileRoute + filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(data)
	if err := w.Error();
	err != nil {
		log.Fatal("error writing csv file", err)
	}
}
