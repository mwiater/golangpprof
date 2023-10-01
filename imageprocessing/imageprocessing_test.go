package imageprocessing

import (
	"bytes"
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testInput        = "../imageprocessing/inputs/input.jpg"
	testOutput       = "../imageprocessing/outputs/grayscaleProcessed.jpg"
	nonExistentImage = "../imageprocessing/inputs/nope.jpg"
)

// isGrayscale checks if the image at the given path is in grayscale.
// It does so by ensuring that for every pixel in the image, the red, green,
// and blue values are equal. If any pixel is found that is not grayscale,
// the function returns false. If an error occurs during file reading or decoding,
// the error is returned.
func isGrayscale(imgPath string) (bool, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return false, err
	}

	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			if r != g || g != b {
				return false, nil
			}
		}
	}

	return true, nil
}

// imagesAreEqual compares two images by their byte sequences to determine
// if they are identical. It decodes the JPEG images from the given paths and
// encodes them back to a byte buffer. The function then compares the byte sequences
// of the two images. If an error occurs during decoding or encoding, the function
// assumes the images are not equal and returns false.
func imagesAreEqual(imgPath1, imgPath2 string) bool {
	img1, err := decodeJPEG(imgPath1)
	if err != nil {
		return false
	}
	img2, err := decodeJPEG(imgPath2)
	if err != nil {
		return false
	}

	buf1 := new(bytes.Buffer)
	err = jpeg.Encode(buf1, img1, nil)
	if err != nil {
		return false
	}

	buf2 := new(bytes.Buffer)
	err = jpeg.Encode(buf2, img2, nil)
	if err != nil {
		return false
	}

	return bytes.Equal(buf1.Bytes(), buf2.Bytes())
}

// Tests for ProcessImageSharpen

// TestProcessImageSharpen_Decode verifies that a valid image can be decoded
// and that an error is returned for a non-existent image.
func TestProcessImageSharpen_Decode(t *testing.T) {
	_, err := decodeJPEG(testInput)
	assert.NoError(t, err)

	_, err = decodeJPEG(nonExistentImage)
	assert.Error(t, err)
}

// TestProcessImageSharpen_Processing ensures that the sharpening process
// completes without errors for a valid input image.
func TestProcessImageSharpen_Processing(t *testing.T) {
	_, err := ProcessImageSharpen(testInput, testOutput)
	assert.NoError(t, err)
}

// TestProcessImageSharpen_OutputSize checks that the output image has a non-zero size.
func TestProcessImageSharpen_OutputSize(t *testing.T) {
	size, err := getFileSize(testOutput)
	assert.NoError(t, err)
	assert.True(t, size > 0)
}

// TestProcessImageSharpen_DifferentFromInput ensures that the sharpened image
// is not identical to the original input image.
func TestProcessImageSharpen_DifferentFromInput(t *testing.T) {
	assert.False(t, imagesAreEqual(testInput, testOutput))
}

// Tests for ProcessImageSharpenOptimized

// TestProcessImageSharpenOptimized_Decode verifies decoding functionality
// for the optimized sharpening function.
func TestProcessImageSharpenOptimized_Decode(t *testing.T) {
	_, err := decodeJPEG(testInput)
	assert.NoError(t, err)

	_, err = decodeJPEG(nonExistentImage)
	assert.Error(t, err)
}

// TestProcessImageSharpenOptimized_Processing tests the optimized sharpening process.
func TestProcessImageSharpenOptimized_Processing(t *testing.T) {
	_, err := ProcessImageSharpenOptimized(testInput, testOutput)
	assert.NoError(t, err)
}

// TestProcessImageSharpenOptimized_OutputSize checks the size of the output
// image for the optimized sharpening function.
func TestProcessImageSharpenOptimized_OutputSize(t *testing.T) {
	size, err := getFileSize(testOutput)
	assert.NoError(t, err)
	assert.True(t, size > 0)
}

// TestProcessImageSharpenOptimized_DifferentFromInput ensures that the output
// from the optimized sharpening function is different from the input.
func TestProcessImageSharpenOptimized_DifferentFromInput(t *testing.T) {
	assert.False(t, imagesAreEqual(testInput, testOutput))
}

// Tests for ProcessImageGrayscale

// TestProcessImageGrayscale_Decode tests the decoding step for the grayscale function
// and ensures an error is returned for a non-existent image.
func TestProcessImageGrayscale_Decode(t *testing.T) {
	_, err := ProcessImageGrayscale(nonExistentImage, testOutput)
	assert.Error(t, err)

	_, err = ProcessImageGrayscale(testInput, testOutput)
	assert.NoError(t, err)
}

// TestProcessImageGrayscale_OutputSize verifies that the grayscale output image
// has a non-zero size.
func TestProcessImageGrayscale_OutputSize(t *testing.T) {
	info, err := os.Stat(testOutput)
	assert.NoError(t, err)
	assert.True(t, info.Size() > 0)
}

// TestProcessImageGrayscale_GrayscaleOutput checks that the processed image is
// indeed in grayscale.
func TestProcessImageGrayscale_GrayscaleOutput(t *testing.T) {
	grayscale, err := isGrayscale(testOutput)
	assert.NoError(t, err)
	assert.True(t, grayscale)
}

// Tests for ProcessImageGrayscaleOptimized

// TestProcessImageGrayscaleOptimized_Decode tests the decoding step for the
// optimized grayscale function.
func TestProcessImageGrayscaleOptimized_Decode(t *testing.T) {
	_, err := ProcessImageGrayscaleOptimized(nonExistentImage, testOutput)
	assert.Error(t, err)

	_, err = ProcessImageGrayscaleOptimized(testInput, testOutput)
	assert.NoError(t, err)
}

// TestProcessImageGrayscaleOptimized_OutputSize verifies the size of the
// output image for the optimized grayscale function.
func TestProcessImageGrayscaleOptimized_OutputSize(t *testing.T) {
	info, err := os.Stat(testOutput)
	assert.NoError(t, err)
	assert.True(t, info.Size() > 0)
}

// TestProcessImageGrayscaleOptimized_GrayscaleOutput ensures that the output
// from the optimized grayscale function is in grayscale.
func TestProcessImageGrayscaleOptimized_GrayscaleOutput(t *testing.T) {
	grayscale, err := isGrayscale(testOutput)
	assert.NoError(t, err)
	assert.True(t, grayscale)
}
