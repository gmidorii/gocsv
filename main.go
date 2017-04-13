package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := flag.String("f", "file.csv", "read file")
	grep := flag.String("g", "", "grep conditions")
	cut := flag.String("c", "1:2", "output field colum num. delimiter :")
	flag.Parse()

	fields := strings.Split(*cut, ":")
	nums, err := strToIntSlice(fields)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}

	csv := csv.NewReader(f)
	records := make([][]string, 0)
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if contains(record, *grep) {
			rs := make([]string, len(nums))
			for i, r := range record {
				for j, num := range nums {
					if i == num {
						rs[j] = r
					}
				}
			}
			records = append(records, rs)
		}
	}

	conv := calc(records)
	fmt.Print(*grep + " " + conv)
}

func contains(iter []string, str string) bool {
	for _, v := range iter {
		if caseInsenstiveContains(v, str) {
			return true
		}
	}
	return false
}

func strToIntSlice(iter []string) ([]int, error) {
	nums := make([]int, len(iter))
	for i, v := range iter {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}

func caseInsenstiveContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}
