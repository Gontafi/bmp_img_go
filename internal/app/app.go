package app

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"bitmap/internal/apply/crop"
	"bitmap/internal/apply/filter"
	mirror "bitmap/internal/apply/miror"
	"bitmap/internal/apply/rotate"
	"bitmap/internal/models"
	"bitmap/internal/read"
	"bitmap/internal/save"
	"bitmap/pkg"
)

var (
	InvalidArg        = errors.New("Invalid args:")
	NotEnoughArgs     = errors.New("Not enough args")
	MissmatchFileType = errors.New("File is not .bmp or .dib type:")
	MissDirection     = errors.New("Miss direction:")
)

func ParseArgsAndRunCommands(args []string) error {
	if len(args) == 1 { //$ ./bitmap case
		pkg.PrintUsage()
		return nil
	}
	if len(args) < 2 {
		return pkg.Check(NotEnoughArgs, "")
	}
	switch args[1] {
	case "apply":

		if pkg.IsHelp(args[2:]) { // .bitmap apply --help Case
			pkg.PrintApplyHelp()
			return nil
		}
		if len(args)-2 <= 2 {
			return pkg.Check(NotEnoughArgs, "")
		}
		imgPath := strings.TrimSpace(args[len(args)-2])
		imgExt := filepath.Ext(imgPath)

		if imgExt != ".bmp" && imgExt != ".dib" {
			return pkg.Check(MissmatchFileType, imgPath)
		}

		file, err := os.Open(imgPath)
		if err != nil {
			return pkg.Check(err, imgPath)
		}

		defer file.Close()

		_, pixels, err := read.ReadImage(file)
		if err != nil {
			return pkg.Check(err, imgPath)
		}

		for _, arg := range args[2 : len(args)-2] {
			parts := strings.Split(arg, "=")
			if len(parts) != 2 {
				return pkg.Check(InvalidArg, "")
			}

			switch parts[0] {
			case "--filter":
				err = filter.ParseFilterDir(pixels, parts[1])
				if err != nil {
					return pkg.Check(err, parts[1])
				}
			case "--rotate":
				pixels, err = ParseRotateDir(pixels, parts[1])
				if err != nil {
					return pkg.Check(err, parts[1])
				}
			case "--mirror":
				pixels, err = ParseMirrorDir(pixels, parts[1])
				if err != nil {
					return pkg.Check(err, parts[1])
				}
			case "--crop":
				pixels, err = ParseCrop(pixels, parts[1])
				if err != nil {
					return pkg.Check(err, parts[1])
				}
			default:
				return pkg.Check(InvalidArg, parts[1])
			}
		}
		err = save.SaveImage(pixels, strings.TrimSpace(args[len(args)-1]))
		if err != nil {
			return err
		}
	case "header":
		if pkg.IsHelp(args) {
			pkg.PrintHeaderHelp()
			return nil
		}
		if len(args) != 3 {
			return NotEnoughArgs
		}

		file, err := os.Open(args[2])
		if err != nil {
			return pkg.Check(err, args[2])
		}

		defer file.Close()

		header, _, err := read.ReadImage(file)
		if err != nil {
			return pkg.Check(err, args[2])
		}
		pkg.PrintHeaderInfo(*header)
	case "--help":
		pkg.PrintUsage()
		return nil
	case "-h":
		pkg.PrintUsage()
		return nil
	default:
		return pkg.Check(InvalidArg, args[1])
	}

	return nil
}

func ParseFilterDir(pixels [][]models.Pixel, s string) {
	panic("unimplemented")
}

func ParseCrop(image [][]models.Pixel, arg string) ([][]models.Pixel, error) {
	sizes := strings.Split(arg, "-")
	switch len(sizes) {
	case 4:
		offsetX, err := strconv.Atoi(sizes[0])
		if err != nil {
			return nil, pkg.Check(err, arg)
		}
		offsetY, err := strconv.Atoi(sizes[1])
		if err != nil {
			return nil, pkg.Check(err, arg)
		}
		width, err := strconv.Atoi(sizes[2])
		if err != nil {
			return nil, pkg.Check(err, arg)
		}
		height, err := strconv.Atoi(sizes[3])
		if err != nil {
			return nil, pkg.Check(err, arg)
		}
		return crop.Crop(image, offsetX, offsetY, width, height)
	case 2:
		offsetX, err := strconv.Atoi(sizes[0])
		if err != nil {
			return nil, pkg.Check(err, arg)
		}
		offsetY, err := strconv.Atoi(sizes[1])
		if err != nil {
			return nil, pkg.Check(err, arg)
		}
		return crop.Crop(image, offsetX, offsetY, 0, 0)
	default:
		return nil, pkg.Check(InvalidArg, arg)
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
		return nil, pkg.Check(MissDirection, direction)
	}
}

func ParseMirrorDir(image [][]models.Pixel, direction string) ([][]models.Pixel, error) {
	switch direction {
	case "horizontal", "h", "hor", "horizontally":
		return mirror.FlipHorizontal(image), nil
	case "vertical", "v", "vertically", "ver":
		return mirror.FlipVertical(image), nil
	default:
		return nil, pkg.Check(MissDirection, direction)
	}
}
