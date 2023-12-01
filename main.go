package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func main() {
	var res int
	readFile, err := os.Open("input")
	if err != nil {
		log.Fatalln(err)
	}
	defer readFile.Close()

	fs := bufio.NewScanner(readFile)
	fs.Split(bufio.ScanLines)

	for fs.Scan() {
		var first, last int
		isFirst := false

		for _, c := range fs.Text() {
			if unicode.IsDigit(c) {
				digit := int(c - '0')
				if !isFirst {
					first = digit
					isFirst = true
				}
				last = digit
			}
		}

		res = res + (first*10 + last)
	}
	print(res)
}