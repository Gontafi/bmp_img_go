package app

import (
	"bitmap/internal/apply/crop"
	mirror "bitmap/internal/apply/miror"
	"bitmap/internal/apply/rotate"
	"bitmap/internal/models"
	"bitmap/internal/read"
	"bitmap/internal/save"
	"bitmap/pkg"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	InvalidArg        = errors.New("invalid args")
	NotEnoughArgs     = errors.New("not enough args")
	MissmatchFileType = errors.New("file is not .bmp or .dib type")
	MissDirection     = errors.New("miss direction")
)

func ParseArgsAndRunCommands(args []string) error {
	if len(args) < 2 {
		return NotEnoughArgs
	}
	switch args[1] {
	case "apply":

		if len(args)-2 <= 2 {
			return NotEnoughArgs
		}
		imgPath := strings.TrimSpace(args[len(args)-2])
		imgExt := filepath.Ext(imgPath)

		if imgExt != ".bmp" && imgExt != ".dib" {
			return MissmatchFileType
		}

		file, err := os.Open(imgPath)
		if err != nil {
			return err
		}

		defer file.Close()

		_, pixels, err := read.ReadImage(file)
		if err != nil {
			return err
		}

		for _, arg := range args[2 : len(args)-2] {
			parts := strings.Split(arg, "=")
			if len(parts) != 2 {
				return InvalidArg
			}

			switch parts[0] {
			case "--filter":
				// to-do
			case "--rotate":
				pixels, err = ParseRotateDir(pixels, parts[1])
				if err != nil {
					return err
				}
			case "--mirror":
				pixels, err = ParseMirrorDir(pixels, parts[1])
				if err != nil {
					return err
				}
			case "--crop":
				pixels, err = ParseCrop(pixels, parts[1])
				if err != nil {
					return err
				}
			default:
				return InvalidArg
			}
		}
		err = save.SaveImage(pixels, strings.TrimSpace(args[len(args)-1]))
		if err != nil {
			return err
		}
	case "header":
		if len(args) != 3 {
			return NotEnoughArgs
		}
		file, err := os.Open(args[2])
		if err != nil {
			return err
		}

		defer file.Close()

		header, _, err := read.ReadImage(file)
		if err != nil {
			return err
		}
		pkg.PrintHeaderInfo(*header)

	default:
		return InvalidArg
	}

	return nil
}

func ParseCrop(image [][]models.Pixel, arg string) ([][]models.Pixel, error) {
	sizes := strings.Split(arg, "-")
	switch len(sizes) {
	case 4:
		offsetX, err := strconv.Atoi(sizes[0])
		if err != nil {
			return nil, err
		}
		offsetY, err := strconv.Atoi(sizes[1])
		if err != nil {
			return nil, err
		}
		width, err := strconv.Atoi(sizes[2])
		if err != nil {
			return nil, err
		}
		height, err := strconv.Atoi(sizes[3])
		if err != nil {
			return nil, err
		}
		return crop.Crop(image, offsetX, offsetY, width, height)
	case 2:
		offsetX, err := strconv.Atoi(sizes[0])
		if err != nil {
			return nil, err
		}
		offsetY, err := strconv.Atoi(sizes[1])
		if err != nil {
			return nil, err
		}
		return crop.Crop(image, offsetX, offsetY, 0, 0)
	default:
		return nil, InvalidArg
	}
}

func ParseRotateDir(image [][]models.Pixel, direction string) ([][]models.Pixel, error) {
	switch direction {
	case "right", "90", "-270":
		return rotate.RotateRight(image), nil
	case "180", "-180":
		return rotate.BottomUp(image), nil
	case "270", "left", "-90":
		return rotate.RotateLeft(image), nil
	case "0", "360", "-360":
		return image, nil
	default:
		break
	}
	return nil, MissDirection
}

func ParseMirrorDir(image [][]models.Pixel, dir string) ([][]models.Pixel, error) {
	switch dir {
	case "horizontal", "h", "hor", "horizontally":
		return mirror.FlipHorizontal(image), nil
	case "vertical", "v", "vertically", "ver":
		return mirror.FlipVertical(image), nil
	default:
		return nil, MissDirection
	}
}
