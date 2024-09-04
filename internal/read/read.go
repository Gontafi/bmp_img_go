package read

import (
	"bitmap/internal/models"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

func ReadHeader(r io.Reader) (*models.BitmapHeader, error) {
	header := make([]byte, 54) // 14 bytes for file header + 40 bytes for DIB header
	_, err := r.Read(header)
	if err != nil {
		return nil, err
	}

	// Parse BMP file header
	fileType := string(header[0:2])
	if fileType != "BM" {
		return nil, fmt.Errorf("not a valid BMP file")
	}

	fileSize := binary.LittleEndian.Uint32(header[2:6])
	// dataOffset := binary.LittleEndian.Uint32(header[10:14])

	// Parse DIB header (BITMAPINFOHEADER)
	headerSize := binary.LittleEndian.Uint32(header[14:18])
	width := binary.LittleEndian.Uint32(header[18:22])
	height := binary.LittleEndian.Uint32(header[22:26])
	planes := binary.LittleEndian.Uint16(header[26:28])
	bitsPerPixel := binary.LittleEndian.Uint16(header[28:30])
	compression := binary.LittleEndian.Uint32(header[30:34])
	imageSize := binary.LittleEndian.Uint32(header[34:38])
	xPixelsPerM := binary.LittleEndian.Uint32(header[38:42])
	yPixelsPerM := binary.LittleEndian.Uint32(header[42:46])
	colorsUsed := binary.LittleEndian.Uint32(header[46:50])
	colorsImportant := binary.LittleEndian.Uint32(header[50:54])

	// Log the extracted values (for debugging purposes)
	log.Printf("FileType: %s", fileType)
	log.Printf("FileSize: %d", fileSize)
	// log.Printf("DataOffset: %d", dataOffset)
	log.Printf("HeaderSize: %d", headerSize)
	log.Printf("Width: %d", width)
	log.Printf("Height: %d", height)
	log.Printf("Planes: %d", planes)
	log.Printf("BitsPerPixel: %d", bitsPerPixel)
	log.Printf("Compression: %d", compression)
	log.Printf("ImageSize: %d", imageSize)
	log.Printf("XPixelsPerM: %d", xPixelsPerM)
	log.Printf("YPixelsPerM: %d", yPixelsPerM)
	log.Printf("ColorsUsed: %d", colorsUsed)
	log.Printf("ColorsImportant: %d", colorsImportant)

	return &models.BitmapHeader{
		FileType:        fileType,
		FileSize:        int(fileSize),
		HeaderSize:      int(headerSize),
		Width:           int(width),
		Height:          int(height),
		Planes:          int(planes),
		BitsPerPixel:    int(bitsPerPixel),
		Compression:     int(compression),
		ImageSize:       int(imageSize),
		XPixelsPerM:     int(xPixelsPerM),
		YPixelsPerM:     int(yPixelsPerM),
		ColorsUsed:      int(colorsUsed),
		ColorsImportant: int(colorsImportant),
	}, nil
}

func ReadImage(r io.Reader, width, height int) ([][]models.Pixel, error) {
	pixels := make([][]models.Pixel, height)
	for i := height - 1; i >= 0; i-- {
		row := make([]models.Pixel, 0)
		for j := 0; j < width; j++ {
			bytes := make([]byte, 3)
			_, err := r.Read(bytes)
			if err != nil {
				return nil, err
			}

			row = append(row, models.Pixel{
				Blue:  bytes[0],
				Green: bytes[1],
				Red:   bytes[2],
			})
		}

		pixels[i] = row
	}

	return pixels, nil
}
