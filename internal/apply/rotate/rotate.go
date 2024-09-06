package rotate

import m "bitmap/internal/models"

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
