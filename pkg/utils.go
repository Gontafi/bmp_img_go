package pkg

import (
	"bitmap/internal/models"
	"fmt"
	"os"
)

func PrintHeaderInfo(header *models.BitmapHeader) {
	fmt.Println("BMP Header:")
	fmt.Println("- FileType ", header.FileType)
	fmt.Println("- FileSizeInBytes ", header.FileSize)
	fmt.Println("- HeaderSize ", header.HeaderSize)
	fmt.Println("DIB Header:")
	fmt.Println("- DibHeaderSize ", header.DibHeaderSize)
	fmt.Println("- WidthInPixels ", header.Width)
	fmt.Println("- HeightInPixels ", header.Height)
	fmt.Println("- PixelSizeInBits ", header.BitsPerPixel)
	fmt.Println("- ImageSizeInBytes ", header.ImageSize)
	fmt.Println("- PixelSize ", header.PixelSize)
	fmt.Println("- Planes ", header.Planes)
	fmt.Println("- Compression ", header.Compression)
	fmt.Println("- XPixelsPerM ", header.XPixelsPerM)
	fmt.Println("- YPixelsPerM ", header.YPixelsPerM)
	fmt.Println("- ColorsUsed ", header.ColorsUsed)
	fmt.Println("- ColorsImportant ", header.ColorsImportant)
}

func Check(err error, text ...string) error {
	fmt.Fprintf(os.Stderr, "Error: %v %s\n\r", err.Error(), text)
	return err
}
