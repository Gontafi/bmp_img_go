package main

import (
	"bitmap/internal/apply/rotate"
	"bitmap/internal/read"
	"bitmap/internal/save"
	"fmt"
	"log"
	"os"
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
	rotated, err := rotate.Rotate(pixels, "90")
	if err == nil {
		fmt.Println(rotated)
	}
	_ = header

	err = save.SaveImage(rotated, "test.bmp")
	if err != nil {
		log.Fatal(err)
	}
}
