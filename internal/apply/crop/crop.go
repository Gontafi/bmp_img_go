package crop

import (
	"errors"

	m "bitmap/internal/models"
)

var ValuesOutOfBound = errors.New("Values out of bound")

func Crop(pixels [][]m.Pixel, offsetX, offsetY, width, heigth int) ([][]m.Pixel, error) {
	if offsetX < 0 || offsetY < 0 || offsetX >= len(pixels[0]) || offsetY >= len(pixels) || offsetX+width >= len(pixels[0]) || offsetY+heigth >= len(pixels) {
		return nil, ValuesOutOfBound
	}
	var cropped [][]m.Pixel

	cropped = pixels[offsetY : offsetY+heigth]
	for i := 0; i < len(cropped); i++ {
		cropped[i] = cropped[i][offsetX : offsetX+width]
	}

	return cropped, nil
}
