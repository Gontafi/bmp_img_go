package main

import (
	"bitmap/internal/read"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.bmp")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = read.ReadHeader(file)
	if err != nil {
		log.Fatal(err)
	}
}
