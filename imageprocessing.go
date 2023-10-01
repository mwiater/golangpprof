package main

import (
	"github.com/mwiater/golangpprof/common"
	"github.com/mwiater/golangpprof/imageprocessing"
)

func main() {
	timedImageProcessGrayscale := common.TimerWrapper(imageprocessing.ProcessImageGrayscale)
	result1 := timedImageProcessGrayscale(imageprocessing.InputPath, imageprocessing.OutputGrayscalePath)

	timedProcessImageGrayscaleOptimized := common.TimerWrapper(imageprocessing.ProcessImageGrayscaleOptimized)
	result2 := timedProcessImageGrayscaleOptimized(imageprocessing.InputPath, imageprocessing.OutputGrayscaleOptimizedPath)

	timedProcessImageSharpen := common.TimerWrapper(imageprocessing.ProcessImageSharpen)
	result3 := timedProcessImageSharpen(imageprocessing.InputPath, imageprocessing.OutputSharpenPath)

	timedProcessImageSharpenOptimized := common.TimerWrapper(imageprocessing.ProcessImageSharpenOptimized)
	result4 := timedProcessImageSharpenOptimized(imageprocessing.InputPath, imageprocessing.OutputSharpenOptimizedPath)

	// Print Results
	//
	//
	common.PrintResults(result1, result2, result3, result4)
}
