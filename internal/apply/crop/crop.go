package crop

import (
	"errors"

	m "bitmap/internal/models"
)

var ValuesOutOfBound = errors.New("Given values out of bound")

// Width and Height may or not be set, if they was not set, then -1 value should be passed to function
func Crop(pixels [][]m.Pixel, offsetX, offsetY, width, heigth int) ([][]m.Pixel, error) {
	if offsetX < 0 || offsetY < 0 || offsetX+width > len(pixels[0]) || offsetY+heigth > len(pixels) || width < 0 || heigth < 0 {
		return nil, ValuesOutOfBound
	}

	// If width and height was not set cut only by offsets
	if width == 0 {
		width = len(pixels[0]) - offsetX
	}
	if heigth == 0 {
		heigth = len(pixels) - offsetY
	}
	cropped := pixels[offsetY : offsetY+heigth] // Cut by Y dimension
	for i := 0; i < len(cropped); i++ {
		cropped[i] = cropped[i][offsetX : offsetX+width] // Cutting by X dimension
	}

	return cropped, nil
}
