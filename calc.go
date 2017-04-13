package main

import (
	"fmt"
	"strconv"
)

func calc(iter [][]string) string {
	// Dependent processing
	var sum int
	var cov int

	for _, v := range iter {
		miss, _ := strconv.Atoi(v[0])
		comp, _ := strconv.Atoi(v[1])
		sum = sum + miss + comp
		cov = cov + comp
	}

	return fmt.Sprintln(strconv.FormatFloat((float64(cov)/float64(sum))*100, 'f', 2, 64))
}
