package imageprocessing

import (
	"image"
	"image/color"
	"sync"
)

var sharpenKernel = [3][3]int{
	{0, -1, 0},
	{-1, 5, -1},
	{0, -1, 0},
}

// See ./imageprocessing/vars.go for defined vars and consts

// ProcessImageSharpen sharpens an image by applying a convolution
// operation using a sharpening kernel. The result is saved to the specified output path.
//
// Parameters:
// - inputPath: Path to the source image which needs to be sharpened.
// - outputPath: Path where the sharpened image will be saved.
//
// Returns:
// - int64: The size of the input image in bytes.
// - error: If any error occurs during the process, it returns the error. Otherwise, it returns nil.
//
// Dependencies:
// - The function relies on external functions and variables:
//   - getFileSize: Retrieves the size of a file.
//   - decodeJPEG: Decodes a JPEG image from a given path.
//   - sharpenKernel: A 3x3 array containing the kernel values for the sharpening operation.
//   - saveProcessedJPEG: Saves the processed image to the given path.
//
// Notes:
// - The image sharpening method used here is a basic convolution with a sharpening kernel.
// - Advanced sharpening techniques might provide better results for specific use cases.
// - The function assumes the image is in JPEG format. If used with another format, it may fail or produce unexpected results.
// - The sharpening kernel values are crucial to the results. A different kernel might produce varied sharpening effects.
func ProcessImageSharpen(inputPath string, outputPath string) (int64, error) {
	size, err := getFileSize(inputPath)
	if err != nil {
		return 0, err
	}
	img, err := decodeJPEG(inputPath)
	if err != nil {
		return 0, err
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	processedImage := image.NewRGBA(bounds)

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			var rSum, gSum, bSum int
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					r, g, b, _ := img.At(x+kx, y+ky).RGBA()
					kernelValue := sharpenKernel[ky+1][kx+1]
					rSum += int(r) * kernelValue
					gSum += int(g) * kernelValue
					bSum += int(b) * kernelValue
				}
			}
			rValue := uint8(min(max(rSum>>8, 0), 255))
			gValue := uint8(min(max(gSum>>8, 0), 255))
			bValue := uint8(min(max(bSum>>8, 0), 255))
			processedImage.Set(x, y, color.RGBA{R: rValue, G: gValue, B: bValue, A: 255})
		}
	}

	err = saveProcessedJPEG(outputPath, processedImage)
	if err != nil {
		return 0, err
	}

	return size, nil
}

// ProcessImageSharpenOptimized sharpens an image by applying a convolution
// operation using a sharpening kernel. The result is saved to the specified output path.
// This function optimizes the sharpening process by dividing the image into sections
// and processing them concurrently with goroutines.
//
// Parameters:
// - inputPath: Path to the source image which needs to be sharpened.
// - outputPath: Path where the sharpened image will be saved.
//
// Returns:
// - int64: The size of the input image in bytes.
// - error: If any error occurs during the process, it returns the error. Otherwise, it returns nil.
//
// Dependencies:
// - The function relies on external functions, variables, and constants:
//   - getFileSize: Retrieves the size of a file.
//   - decodeJPEG: Decodes a JPEG image from a given path.
//   - sharpenKernel: A 3x3 array containing the kernel values for the sharpening operation.
//   - saveProcessedJPEG: Saves the processed image to the given path.
//   - NumRoutines: Constant that dictates how many goroutines should be spawned for concurrent processing.
//
// Notes:
// - This optimized function breaks the image into sections and processes them concurrently for faster results.
// - The sharpening kernel values remain crucial to the results. A different kernel might produce varied sharpening effects.
func ProcessImageSharpenOptimized(inputPath string, outputPath string) (int64, error) {
	size, err := getFileSize(inputPath)
	if err != nil {
		return 0, err
	}
	img, err := decodeJPEG(inputPath)
	if err != nil {
		return 0, err
	}

	bounds := img.Bounds()
	processedImage := image.NewRGBA(bounds)

	var wg sync.WaitGroup
	step := bounds.Dy() / NumRoutines

	for i := 0; i < NumRoutines; i++ {
		startY := i * step
		endY := (i + 1) * step
		if i == NumRoutines-1 {
			endY = bounds.Max.Y
		}
		wg.Add(1)
		go sharpenConcurrent(img, startY, endY, &wg, processedImage)
	}
	wg.Wait()

	err = saveProcessedJPEG(outputPath, processedImage)
	if err != nil {
		return 0, err
	}

	return size, nil
}

// sharpenConcurrent is a helper function for ProcessImageSharpenOptimized.
// It processes a section of the image by applying the sharpening convolution
// from start to end rows concurrently.
//
// Parameters:
// - img: The original image that needs sharpening.
// - start: Starting row of the section to be processed.
// - end: Ending row of the section to be processed.
// - wg: WaitGroup to signal when the goroutine has finished its work.
// - output: Image to store the sharpened result.
//
// Dependencies:
// - The function uses the sharpenKernel for convolution.
//
// Notes:
// - This function is intended to be used as a goroutine.
// - The function doesn't handle border rows since they don't have enough neighbors for convolution.
func sharpenConcurrent(img image.Image, start, end int, wg *sync.WaitGroup, output *image.RGBA) {
	defer wg.Done()
	bounds := img.Bounds()
	width := bounds.Dx()

	for y := start; y < end; y++ {
		for x := 1; x < width-1; x++ {
			var rSum, gSum, bSum int
			for ky := -1; ky <= 1; ky++ {
				for kx := -1; kx <= 1; kx++ {
					r, g, b, _ := img.At(x+kx, y+ky).RGBA()
					kernelValue := sharpenKernel[ky+1][kx+1]
					rSum += int(r) * kernelValue
					gSum += int(g) * kernelValue
					bSum += int(b) * kernelValue
				}
			}
			rValue := uint8(min(max(rSum>>8, 0), 255))
			gValue := uint8(min(max(gSum>>8, 0), 255))
			bValue := uint8(min(max(bSum>>8, 0), 255))
			output.Set(x, y, color.RGBA{R: rValue, G: gValue, B: bValue, A: 255})
		}
	}
}
