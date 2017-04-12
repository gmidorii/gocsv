package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file := flag.String("f", "files.csv", "read file")
	flag.Parse()
	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	csv := csv.NewReader(f)
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
