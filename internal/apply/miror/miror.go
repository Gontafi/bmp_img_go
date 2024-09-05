package mirror

import "bitmap/internal/models"

func FlipHorizontal(pixels [][]models.Pixel) [][]models.Pixel {
	height := len(pixels)
	if height == 0 {
		return pixels
	}
	width := len(pixels[0])

	flipped := make([][]models.Pixel, height)
	for y := 0; y < height; y++ {
		flipped[y] = make([]models.Pixel, width)
		for x := 0; x < width; x++ {
			flipped[y][x] = pixels[y][width-x-1]
		}
	}
	return flipped
}

func FlipVertical(pixels [][]models.Pixel) [][]models.Pixel {
	height := len(pixels)
	if height == 0 {
		return pixels
	}

	flipped := make([][]models.Pixel, height)
	for y := 0; y < height; y++ {
		flipped[y] = pixels[height-y-1]
	}
	return flipped
}
