package rotate

import (
	m "bitmap/internal/models"
	"fmt"
)

func Rotate(image [][]m.Pixel, direction string) ([][]m.Pixel, error) {
	if len(image) == 0 {
		return nil, fmt.Errorf("Image lenght error")
	}
	switch direction {
	case "right", "90", "-270":
		return RotateRight(image), nil
	case "180", "-180":
		return BottomUp(image), nil
	case "270", "left", "-90":
		return RotateLeft(image), nil
	case "0", "360", "-360":
		return image, nil
	default:
		break
	}
	return nil, fmt.Errorf("Error direction")
}

func RotateRight(image [][]m.Pixel) [][]m.Pixel {
	var rotatedImage [][]m.Pixel
	for i := 0; i < len(image[0]); i++ {
		row := []m.Pixel{}
		for j := 0; j < len(image); j++ {
			pixel := image[len(image)-1-j][i]
			row = append(row, pixel)
		}
		rotatedImage = append(rotatedImage, row)
	}
	return rotatedImage
}

func RotateLeft(image [][]m.Pixel) [][]m.Pixel {
	var rotatedImage [][]m.Pixel
	for i := len(image[0]) - 1; i >= 0; i-- {
		row := []m.Pixel{}
		for j := 0; j < len(image); j++ {
			pixel := image[j][i]
			row = append(row, pixel)
		}
		rotatedImage = append(rotatedImage, row)
	}
	return rotatedImage
}

func BottomUp(image [][]m.Pixel) [][]m.Pixel {
	var rotatedImage [][]m.Pixel
	for i := len(image) - 1; i >= 0; i-- {
		row := []m.Pixel{}
		for j := len(image[0]) - 1; j >= 0; j-- {
			pixel := image[i][j]
			row = append(row, pixel)
		}
		rotatedImage = append(rotatedImage, row)
	}
	return rotatedImage
}
