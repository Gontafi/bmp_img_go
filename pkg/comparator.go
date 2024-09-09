package pkg

import (
	"bitmap/internal/models"
	"fmt"
)

func Compare(pix1 [][]models.Pixel, pix2 [][]models.Pixel) {
	for i := 0; i < len(pix1); i++ {
		for j := 0; j < len(pix1[i]); j++ {
			fmt.Println("Green difference:", pix1[i][j].Green-pix2[i][j].Green)
			fmt.Println("Red difference:", pix1[i][j].Red-pix2[i][j].Red)
		}
	}
}
