package main

import (
	"bitmap/internal/read"
	"bitmap/internal/save"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample.bmp")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	header, err := read.ReadHeader(file)
	if err != nil {
		log.Fatal(err)
	}

	pixels, err := read.ReadImage(file, header.Width, header.Height)
	if err != nil {
		log.Fatal(err)
	}

	err = save.SaveImage(pixels, "test.bmp")
	if err != nil {
		log.Fatal(err)
	}
}
