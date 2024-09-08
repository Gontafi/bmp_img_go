package pkg

import "fmt"

func printUsage() {
	usageText := `Usage:
  bitmap <command> [arguments]

The commands are:
  header    prints bitmap file header information
  apply     applies processing to the image and saves it to the file
`
	fmt.Print(usageText)
}

func printHeaderHelp() {
	headerHelpText := `Usage:
  bitmap header <source_file>

Description:
  Prints bitmap file header information
`
	fmt.Print(headerHelpText)
}

func printApplyHelp() {
	applyHelpText := `Usage:
  bitmap apply [options] <source_file> <output_file>

The options are:
  -h, --help      																			    
  		Prints program usage information
  --mirror=Value	Values: horizontal, h, horizontally, hor, vertical, v, vertically, ver      
  		Mirroring a photo vertically is replacing pixels from top to bottom.
  --filter=Value	Values: blue, red, green, grayscale, negative, pixelate, blur			    
  		Applies given effect to the image.
  --rotate=Value    Values: right, left  														
  		Rotates image for 90 degrees to the given direction
  --crop=OffsetX-OffsetY-Width-Height 															
  		Accepts values that specify the offset by X, the offset by Y, the width, and the height. (--crop=OffsetX-OffsetY-Width-Height). Where width and height are optional.    
	
	These option can be used multiple times.
`
	fmt.Print(applyHelpText)
}
