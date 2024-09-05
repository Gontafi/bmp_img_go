package pkg

import (
	"fmt"
	"os"

	"bitmap/internal/models"
)

func PrintHeaderInfo(header models.BitmapHeader) {
	fmt.Println("BMP Header:")
	fmt.Println("- FileType ", header.FileType)
	fmt.Println("- FileSizeInBytes ", header.FileSize)
	fmt.Println("- HeaderSize ", header.HeaderSize)
	fmt.Println("DIB Header:")
	fmt.Println("- DibHeaderSize ", header.DibHeaderSize)
	fmt.Println("- WidthInPixels ", header.Width)
	fmt.Println("- HeightInPixels ", header.Height)
	fmt.Println("- PixelSize ", header.PixelSize)
	fmt.Println("- ImageSize ", header.ImageSize)
	fmt.Println("- Planes ", header.Planes)
	fmt.Println("- BitsPerPixel ", header.BitsPerPixel)
	fmt.Println("- Compression ", header.Compression)
	fmt.Println("- XPixelsPerM ", header.XPixelsPerM)
	fmt.Println("- YPixelsPerM ", header.YPixelsPerM)
	fmt.Println("- ColorsUsed ", header.ColorsUsed)
	fmt.Println("- ColorsImportant ", header.ColorsImportant)
}

func Check(err error, text string) error {
	fmt.Fprintf(os.Stderr, "Error: %v %s\n\r", err.Error(), text)
	return err
}
