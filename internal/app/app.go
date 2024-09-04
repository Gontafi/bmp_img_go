package app

import (
	"bitmap/internal/apply/rotate"
	"bitmap/internal/read"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseArgsAndRunCommands(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough args")
	}

	switch args[2] {
	case "apply":

		if len(args)-2 <= 2 {
			return fmt.Errorf("not enough args")
		}
		imgPath := strings.TrimSpace(args[len(args)-1])
		imgExt := filepath.Ext(imgPath)

		if imgExt != ".bmp" && imgExt != ".dib" {
			return fmt.Errorf("file is not bmp")
		}

		file, err := os.Open(imgPath)
		if err != nil {
			return err
		}

		defer file.Close()

		for _, arg := range args[2 : len(args)-2] {
			parts := strings.Split(arg, "=")
			if len(parts) != 2 {
				return fmt.Errorf("invalid args")
			}

			switch parts[0] {
			case "filter":
				// to-do
			case "rotate":
				rotate.R
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

		header, _, err := read.ReadImage(file)
		if err != nil {
			return err
		}

	}
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
