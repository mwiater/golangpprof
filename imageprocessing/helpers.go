package imageprocessing

import (
	"image"
	"image/jpeg"
	"os"
)

// getFileSize retrieves the size of the file located at the specified path.
//
// Parameters:
// - inputPath: Path to the target file.
//
// Returns:
// - int64: The size of the file in bytes.
// - error: If any error occurs during retrieval, it returns the error. Otherwise, it returns nil.
func getFileSize(inputPath string) (int64, error) {
	fi, err := os.Stat(inputPath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

// decodeJPEG decodes a JPEG image located at the specified path.
//
// Parameters:
// - inputPath: Path to the source JPEG image.
//
// Returns:
// - image.Image: The decoded image.
// - error: If any error occurs during decoding, it returns the error. Otherwise, it returns nil.
func decodeJPEG(inputPath string) (image.Image, error) {
	input, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	img, err := jpeg.Decode(input)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// saveProcessedJPEG saves the given RGBA image as a JPEG to the specified path.
//
// Parameters:
// - outputPath: Path where the image will be saved.
// - processedImage: The RGBA image to be saved.
//
// Returns:
// - error: If any error occurs during saving, it returns the error. Otherwise, it returns nil.
func saveProcessedJPEG(outputPath string, processedImage *image.RGBA) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, processedImage, nil)
	if err != nil {
		return err
	}

	return nil
}

// saveProcessedGrayScaleJPEG saves the given grayscale image as a JPEG to the specified path.
//
// Parameters:
// - outputPath: Path where the grayscale image will be saved.
// - processedImage: The grayscale image to be saved.
//
// Returns:
// - error: If any error occurs during saving, it returns the error. Otherwise, it returns nil.
func saveProcessedGrayScaleJPEG(outputPath string, processedImage *image.Gray) error {
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = jpeg.Encode(outputFile, processedImage, nil)
	if err != nil {
		return err
	}

	return nil
}

// min returns the minimum of the two provided integers.
//
// Parameters:
// - a, b: The integers to compare.
//
// Returns:
// - int: The smaller of the two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of the two provided integers.
//
// Parameters:
// - a, b: The integers to compare.
//
// Returns:
// - int: The larger of the two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
