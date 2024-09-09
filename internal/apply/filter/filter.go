package filter

import (
	"bitmap/internal/models"
	"bitmap/pkg"
	"errors"
)

var ErrInvalidArgument error = errors.New("Invalid argument passed to flag:")

func ParseFilterDir(pix [][]models.Pixel, argument string) error {
	switch argument {
	case "blue":
		blue(pix)
	case "red":
		red(pix)
	case "green":
		green(pix)
	case "grayscale":
		grayscale(pix)
	case "negative":
		negative(pix)
	case "pixelate":
		pixelate(pix)
	case "blur":
		blur(pix)
	default:
		return ErrInvalidArgument
	}
	return nil
}

func grayscale(pix [][]models.Pixel) {
	for i := 0; i < len(pix); i++ {
		for j := 0; j < len(pix[i]); j++ {
			average := (pix[i][j].Blue + pix[i][j].Red + pix[i][j].Green) / 3
			pix[i][j].Blue = average
			pix[i][j].Green = average
			pix[i][j].Red = average
		}
	}
	return
}

func blue(pix [][]models.Pixel) {
	for i := 0; i < len(pix); i++ {
		for j := 0; j < len(pix[i]); j++ {
			pix[i][j].Green = 0
			pix[i][j].Red = 0
		}
	}
	return
}

func green(pix [][]models.Pixel) {
	for i := 0; i < len(pix); i++ {
		for j := 0; j < len(pix[i]); j++ {
			pix[i][j].Blue = 0
			pix[i][j].Red = 0
		}
	}
	return
}

func red(pix [][]models.Pixel) {
	for i := 0; i < len(pix); i++ {
		for j := 0; j < len(pix[i]); j++ {
			pix[i][j].Blue = 0
			pix[i][j].Green = 0
		}
	}
	return
}

func negative(pix [][]models.Pixel) {
	for i := 0; i < len(pix); i++ {
		for j := 0; j < len(pix[i]); j++ {
			pix[i][j].Red = 255 - pix[i][j].Red
			pix[i][j].Blue = 255 - pix[i][j].Blue
			pix[i][j].Green = 255 - pix[i][j].Green
		}
	}
	return
}

func pixelate(pix [][]models.Pixel) {
	for i := 19; i < len(pix); i += 20 {
		for j := 19; j < len(pix[i]); j += 20 {
			averageB := 0
			averageG := 0
			averageR := 0
			for k := i - 19; k <= i; k++ {
				for m := j - 19; m <= j; m++ {
					averageB += int(pix[k][m].Blue)
					averageG += int(pix[k][m].Green)
					averageR += int(pix[k][m].Red)
				}
			}
			averageB /= 400
			averageG /= 400
			averageR /= 400
			for k := i - 19; k <= i; k++ {
				for m := j - 19; m <= j; m++ {
					pix[k][m].Blue = byte(averageB)
					pix[k][m].Green = byte(averageG)
					pix[k][m].Red = byte(averageR)
				}
			}

		}
	}
	return
}

func blur(pix [][]models.Pixel) {
	neigbors := pkg.GenerateNeighbours(10)

	for i := 0; i < len(pix); i++ {
		for j := 0; j < len(pix[i]); j++ {

			averageB := 0
			averageG := 0
			averageR := 0

			counter := 0
			for _, k := range neigbors {
				newX := i + k[0]
				newY := j + k[1]
				if newX >= 0 && newX < len(pix) && newY >= 0 && newY < len(pix[i]) {
					averageB += int(pix[newX][newY].Blue)
					averageG += int(pix[newX][newY].Green)
					averageR += int(pix[newX][newY].Red)
					counter++
				}
			}

			averageB /= counter
			averageG /= counter
			averageR /= counter

			pix[i][j].Blue = byte(averageB)
			pix[i][j].Red = byte(averageR)
			pix[i][j].Green = byte(averageG)
		}
	}
	return
}
