package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	data, err := os.ReadFile("test.lang")
	if err != nil {
		panic(err.Error())
	}

	parsedData := string(data)
	result := strings.ReplaceAll(parsedData, " ", "")

	for _, r := range result {
		if string(r) == "+" {

		}
		fmt.Println(string(r))
	}
}
