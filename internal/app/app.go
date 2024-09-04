package app

import (
	"bitmap/internal/read"
	"fmt"
	"os"
	"strings"
)

func ParseArgsAndRunCommands(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough args")
	}

	switch args[2] {
	case "apply":
		for _, arg := range args[2:] {
			parts := strings.Split(arg, "=")
			if len(parts) != 2 {
				return fmt.Errorf("invalid args")
			}

			switch parts[0] {
			case "filter":
				// to-do
			case "rotate":
				// to-do
			case "mirror":
				// to-do
			case "crop":
				// to-do
			default:
				fmt.Println("tut error kakoi to")
			}
		}
	case "header":
		if len(args) != 4 {
			fmt.Println("tut error kakoi to")
		}
		file, err := os.Open(args[3])
		if err != nil {
			return err
		}

		defer file.Close()

		_, _, err = read.ReadImage(file)
		if err != nil {
			return err
		}

	}

	return nil
}

/*
$ ./bitmap header sample.bmp
BMP Header:
- FileType BM
- FileSizeInBytes 518456
- HeaderSize 54
DIB Header:
- DibHeaderSize 40
- WidthInPixels 480
- HeightInPixels 360
- PixelSizeInBits 24
- ImageSizeInBytes 518402
$
*/
