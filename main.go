package main

import (
	"log"
	"os"

	"bitmap/internal/read"
	"bitmap/internal/save"
)

func main() {
	file, err := os.Open("sample.bmp")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	header, pixels, err := read.ReadImage(file)
	if err != nil {
		log.Fatal(err)
	}

	_ = header

	err = save.SaveImage(pixels, "test.bmp")
	if err != nil {
		log.Fatal(err)
	}
}
