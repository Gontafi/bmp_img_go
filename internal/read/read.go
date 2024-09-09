package read

import (
	"encoding/binary"
	"fmt"
	"io"

	"bitmap/internal/models"
)

func ReadImage(r io.ReadSeeker) (*models.BitmapHeader, [][]models.Pixel, error) {
	header := make([]byte, 54) // 14 bytes for file header + 40 bytes for DIB header
	_, err := r.Read(header)
	if err != nil {
		return nil, nil, err
	}

	// Parse BMP file header
	fileType := string(header[0:2])
	if fileType != "BM" {
		return nil, nil, fmt.Errorf("not a valid BMP file")
	}

	fileSize := binary.LittleEndian.Uint32(header[2:6])
	dataOffset := binary.LittleEndian.Uint32(header[10:14])

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

	if compression != 0 {
		return nil, nil, fmt.Errorf("unsupported BMP compression type")
	}

	// Move the reader to the start of the image data
	_, err = r.Seek(int64(dataOffset), io.SeekStart)
	if err != nil {
		return nil, nil, err
	}

	// Calculate row padding
	rowSize := (int(width)*int(bitsPerPixel)/8 + 3) &^ 3
	padding := rowSize - int(width)*int(bitsPerPixel)/8

	pixels := make([][]models.Pixel, height)
	for i := int(height) - 1; i >= 0; i-- {
		row := make([]models.Pixel, width)
		for j := 0; j < int(width); j++ {
			bytes := make([]byte, 3)
			_, err := r.Read(bytes)
			if err != nil {
				return nil, nil, err
			}

			row[j] = models.Pixel{
				Blue:  bytes[0],
				Green: bytes[1],
				Red:   bytes[2],
			}
		}
		// Skip the padding bytes
		_, err = r.Seek(int64(padding), io.SeekCurrent)
		if err != nil {
			return nil, nil, err
		}

		pixels[i] = row
	}

	return &models.BitmapHeader{
		FileType:        fileType,
		FileSize:        int(fileSize),
		HeaderSize:      int(dataOffset),
		DibHeaderSize:   int(headerSize),
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
	}, pixels, nil
}
