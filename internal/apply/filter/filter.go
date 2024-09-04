package filter

import "bitmap/internal/models"

func Filter(pix [][]models.Pixel) error {
	blur(pix)
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
	for i := 2; i < len(pix); i += 3 {
		for j := 2; j < len(pix[i]); j += 3 {
			averageB := 0
			averageG := 0
			averageR := 0
			for k := i - 2; k <= i; k++ {
				for m := j - 2; m <= j; m++ {
					averageB += int(pix[k][m].Blue)
					averageG += int(pix[k][m].Green)
					averageR += int(pix[k][m].Red)
				}
			}
			averageB /= 9
			averageG /= 9
			averageR /= 9
			for k := i - 2; k <= i; k++ {
				for m := j - 2; m <= j; m++ {
					pix[k][m].Blue = byte(averageB)
					pix[k][m].Green = byte(averageG)
					pix[k][m].Red = byte(averageR)
				}
			}

		}
	}
	return
}
