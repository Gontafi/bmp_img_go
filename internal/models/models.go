package models

// BitmapHeader represents the metadata information of a bitmap file.
type BitmapHeader struct {
	FileType        string // BM
	FileSize        int    // File size in bytes
	HeaderSize      int    // Header size in bytes
	DibHeaderSize   int    // DIB header size in bytes
	Width           int    // Width of the image in pixels
	Height          int    // Height of the image in pixels
	PixelSize       int    // Bits per pixel
	ImageSize       int    // Image size in bytes
	Planes          int
	BitsPerPixel    int
	Compression     int
	XPixelsPerM     int
	YPixelsPerM     int
	ColorsUsed      int
	ColorsImportant int
}

// CropParams represents the cropping parameters for the bitmap image.
type CropParams struct {
	OffsetX int // X offset for cropping
	OffsetY int // Y offset for cropping
	Width   int // Optional width for cropping
	Height  int // Optional height for cropping
}

type Pixel struct {
	Blue  byte
	Green byte
	Red   byte
}

type ImageBody struct {
	Width  int
	Height int
	Pixels [][]Pixel
}
