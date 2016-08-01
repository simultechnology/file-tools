package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	fmt.Println("start!")
	files := os.Args[1:]

	fmt.Printf("%v\n", files)

	var filters []string
	var referers []string

	if len(files) > 1 {

		filterPointer, _ := os.Open(files[0])
		defer filterPointer.Close()

		filterInput := bufio.NewScanner(filterPointer)
		for filterInput.Scan() {
			filters = append(filters, filterInput.Text())
		}

		refererPointer, _ := os.Open(files[1])
		defer refererPointer.Close()

		refererInput := bufio.NewScanner(refererPointer)
		for refererInput.Scan() {
			referers = append(referers, refererInput.Text())
		}

	}

	fmt.Printf("%v\n", filters)
	fmt.Printf("%v\n", referers)

	for _, referer := range referers {

		contains := false

		for _, filter := range filters {
			splits := strings.Split(referer, ",")
			contains = strings.HasPrefix(splits[1], filter)
			if contains {
				break
			}
		}
		if contains {
			fmt.Printf("%s\n", referer)
		}
	}

}
