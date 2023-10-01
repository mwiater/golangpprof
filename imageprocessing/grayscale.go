package imageprocessing

import (
	"image"
	"image/color"
	"sync"
)

// See ./imageprocessing/vars.go for defined vars and consts

// ProcessImageGrayscale converts an image to grayscale and saves the result to the specified output path.
//
// Parameters:
// - inputPath: Path to the source image which needs to be converted to grayscale.
// - outputPath: Path where the grayscale image will be saved.
//
// Returns:
// - int64: The size of the input image in bytes.
// - error: If any error occurs during the process, it returns the error. Otherwise, it returns nil.
//
// Dependencies:
// - The function relies on external functions:
//   - getFileSize: Retrieves the size of a file.
//   - decodeJPEG: Decodes a JPEG image from a given path.
//   - saveProcessedGrayScaleJPEG: Saves the grayscale processed image to the given path.
//
// Notes:
// - The function assumes the image is in JPEG format. If used with another format, it may fail or produce unexpected results.
// - The conversion method retrieves the red channel from the original color, assuming that it's representative of the grayscale.
//   More sophisticated methods for grayscale conversion might consider weighted averages of RGB values.
func ProcessImageGrayscale(inputPath string, outputPath string) (int64, error) {
	size, err := getFileSize(inputPath)
	if err != nil {
		return 0, err
	}
	img, err := decodeJPEG(inputPath)
	if err != nil {
		return 0, err
	}

	bounds := img.Bounds()
	processedImage := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			gray, _, _, _ := c.RGBA()
			processedImage.Set(x, y, color.Gray{Y: uint8(gray >> 8)})
		}
	}

	err = saveProcessedGrayScaleJPEG(outputPath, processedImage)
	if err != nil {
		return 0, err
	}

	return size, nil
}

// ProcessImageGrayscaleOptimized converts an image to grayscale using concurrency
// and saves the result to the specified output path. This optimized function
// splits the image into sections and processes them concurrently for faster results.
//
// Parameters:
// - inputPath: Path to the source image which needs to be converted to grayscale.
// - outputPath: Path where the grayscale image will be saved.
//
// Returns:
// - int64: The size of the input image in bytes.
// - error: If any error occurs during the process, it returns the error. Otherwise, it returns nil.
//
// Dependencies:
// - The function relies on external functions, variables, and constants:
//   - getFileSize: Retrieves the size of a file.
//   - decodeJPEG: Decodes a JPEG image from a given path.
//   - saveProcessedGrayScaleJPEG: Saves the grayscale processed image to the given path.
//   - NumRoutines: Constant that dictates how many goroutines should be spawned for concurrent processing.
//
// Notes:
// - This optimized function breaks the image into sections and processes them concurrently for faster results.
// - The conversion method retrieves the red channel from the original color,
//   assuming that it's representative of the grayscale. Advanced grayscale conversion might consider weighted averages of RGB values.
func ProcessImageGrayscaleOptimized(inputPath string, outputPath string) (int64, error) {
	size, err := getFileSize(inputPath)
	if err != nil {
		return 0, err
	}
	img, err := decodeJPEG(inputPath)
	if err != nil {
		return 0, err
	}

	bounds := img.Bounds()
	processedImage := image.NewGray(bounds)

	var wg sync.WaitGroup
	step := bounds.Dy() / NumRoutines

	for i := 0; i < NumRoutines; i++ {
		startY := i * step
		endY := (i + 1) * step
		if i == NumRoutines-1 {
			endY = bounds.Max.Y
		}
		wg.Add(1)
		go toGrayscaleConcurrent(img, startY, endY, &wg, processedImage)
	}
	wg.Wait()

	err = saveProcessedGrayScaleJPEG(outputPath, processedImage)
	if err != nil {
		return 0, err
	}

	return size, nil
}

// toGrayscaleConcurrent is a helper function for ProcessImageGrayscaleOptimized.
// It processes a section of the image to convert it to grayscale concurrently.
//
// Parameters:
// - img: The original image that needs to be converted.
// - start: Starting row of the section to be processed.
// - end: Ending row of the section to be processed.
// - wg: WaitGroup to signal when the goroutine has finished its work.
// - grayImage: Image to store the grayscale result.
//
// Notes:
// - This function is intended to be used as a goroutine.
func toGrayscaleConcurrent(img image.Image, start, end int, wg *sync.WaitGroup, grayImage *image.Gray) {
	defer wg.Done()

	bounds := img.Bounds()
	for y := start; y < end; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.At(x, y)
			gray, _, _, _ := c.RGBA()
			grayImage.Set(x, y, color.Gray{Y: uint8(gray >> 8)})
		}
	}
}
