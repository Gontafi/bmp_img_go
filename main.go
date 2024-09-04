package main

import (
	"log"
	"os"

	"bitmap/internal/apply/filter"
	"bitmap/internal/read"
	"bitmap/internal/save"
)

func main() {
	file, err := os.Open("sample.bmp")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, pixels, err := read.ReadImage(file)
	if err != nil {
		log.Fatal(err)
	}

	filter.Filter(pixels)

	err = save.SaveImage(rotated, "test.bmp")
	if err != nil {
		log.Fatal(err)
	}
}
