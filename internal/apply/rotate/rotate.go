package rotate

import (
	m "bitmap/internal/models"
	"fmt"
)

func Rotate(image [][]m.Pixel, direction string) ([][]m.Pixel, error) {
	rotated := [][]m.Pixel{}
	var turn int
	switch direction {
	case "right", "90", "-270":
		if len(image) != 0 {
			rotated = make([][]m.Pixel, len(image[0]))
			for i := 0; i < len(image); i++ {
				rotated[i] = make([]m.Pixel, len(image))
			}
			turn = 1
		}
	case "180", "-180":
		turn = 2
	case "270", "left", "-90":
		if len(image) != 0 {
			rotated = make([][]m.Pixel, len(image[0]))
			for i := 0; i < len(image); i++ {
				rotated[i] = make([]m.Pixel, len(image))
			}
			turn = 1
		}
		turn = 3
	default:
		turn = 0
	}
	if turn == 0 {
		return nil, fmt.Errorf("Direction error")
	}
	for t := 0; t < turn; t++ {
		for i := 0; i < len(image); i++ {
			row := []m.Pixel{}
			for j := 0; j < len(image[i]); j++ {
				pixel := image[len(image)-1-j][i]

				row = append(row, pixel)
			}
			rotated = append(rotated, row)
		}
	}

	return rotated, nil
}

/*func Rotate(image [][]int, direction string) ([][]int, error) {
	rotated := [][]int{}
	var turn int
	switch direction {
	case "right", "90", "-270":
		if len(image) != 0 {
			rotated = make([][]int, len(image[0]))
			for i := 0; i < len(image[i]); i++ {
				rotated[i] = make([]int, len(image))
			}
			turn = 1
		}
	case "180", "-180":
		turn = 2
	case "270", "left", "-90":
		if len(image) != 0 {
			rotated = make([][]int, len(image[0]))
			for i := 0; i < len(image[i]); i++ {
				rotated[i] = make([]int, len(image))
			}
			turn = 1
		}
		turn = 3
	default:
		turn = 0
	}
	if turn == 0 {
		return nil, fmt.Errorf("Direction error")
	}

	for t := 0; t < turn; t++ {
		for i := 0; i < len(image[0]); i++ {
			row := []int{}
			for j := 0; j < len(image); j++ {
				pixel := image[len(image)-1-j][i]
				fmt.Println(pixel)
				row = append(row, pixel)
			}
			rotated = append(rotated, row)
		}
	}

	return rotated, nil
}
*/
