package save

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	"bitmap/internal/models"
)

// Calculate the padding needed for each row to be aligned to a 4-byte boundary
func calculateRowPadding(width int) int {
	rowSize := width * 3 // 3 bytes per pixel (RGB)
	padding := (4 - (rowSize % 4)) % 4
	return padding
}

// SaveImage writes the pixel data to a BMP file
func SaveImage(pixels [][]models.Pixel, fileName string) error {
	height := len(pixels)
	if height == 0 {
		return fmt.Errorf("pixels data is empty")
	}
	width := len(pixels[0])

	rowPadding := calculateRowPadding(width)
	rowSize := width*3 + rowPadding

	fileSize := 14 + 40 + height*rowSize // File header (14 bytes) + DIB header (40 bytes) + pixel data

	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the file header
	fileHeader := bytes.Buffer{}
	fileHeader.Write([]byte{'B', 'M'}) // File type
	binary.Write(&fileHeader, binary.LittleEndian, uint32(fileSize))
	binary.Write(&fileHeader, binary.LittleEndian, uint16(0))     // Reserved1
	binary.Write(&fileHeader, binary.LittleEndian, uint16(0))     // Reserved2
	binary.Write(&fileHeader, binary.LittleEndian, uint32(14+40)) // Offset to pixel data

	// Write the DIB header (BITMAPINFOHEADER)
	dibHeader := bytes.Buffer{}
	binary.Write(&dibHeader, binary.LittleEndian, uint32(40))    // DIB header size
	binary.Write(&dibHeader, binary.LittleEndian, int32(width))  // Image width
	binary.Write(&dibHeader, binary.LittleEndian, int32(height)) // Image height
	binary.Write(&dibHeader, binary.LittleEndian, uint16(1))     // Number of color planes
	binary.Write(&dibHeader, binary.LittleEndian, uint16(24))    // Bits per pixel
	binary.Write(&dibHeader, binary.LittleEndian, uint32(0))     // Compression (0 = none)
	binary.Write(&dibHeader, binary.LittleEndian, uint32(0))     // Image size (can be 0 for uncompressed)
	binary.Write(&dibHeader, binary.LittleEndian, uint32(2835))  // Horizontal resolution (pixels per meter)
	binary.Write(&dibHeader, binary.LittleEndian, uint32(2835))  // Vertical resolution (pixels per meter)
	binary.Write(&dibHeader, binary.LittleEndian, uint32(0))     // Colors used
	binary.Write(&dibHeader, binary.LittleEndian, uint32(0))     // Important colors

	// Write headers to file
	file.Write(fileHeader.Bytes())
	file.Write(dibHeader.Bytes())

	// Write pixel data
	for i := height - 1; i >= 0; i-- {
		row := pixels[i]
		for j := 0; j < width; j++ {
			pixel := row[j]
			file.Write([]byte{pixel.Blue, pixel.Green, pixel.Red})
		}
		// Write padding bytes
		if rowPadding > 0 {
			file.Write(make([]byte, rowPadding))
		}
	}

	return nil
}
