package main

import (
	"github.com/mwiater/golangpprof/common"
	"github.com/mwiater/golangpprof/imageprocessing"
)

func main() {
	// Pass ProcessImageGrayscale() to TimerWrapper()
	timedProcessImageGrayscale := common.TimerWrapper(imageprocessing.ProcessImageGrayscale)
	result1 := timedProcessImageGrayscale(imageprocessing.InputPath, imageprocessing.OutputGrayscalePath)

	// Pass ProcessImageGrayscaleOptimized() to TimerWrapper()
	timedProcessImageGrayscaleOptimized := common.TimerWrapper(imageprocessing.ProcessImageGrayscaleOptimized)
	result2 := timedProcessImageGrayscaleOptimized(imageprocessing.InputPath, imageprocessing.OutputGrayscaleOptimizedPath)

	// Pass ProcessImageSharpen() to TimerWrapper()
	timedProcessImageSharpen := common.TimerWrapper(imageprocessing.ProcessImageSharpen)
	result3 := timedProcessImageSharpen(imageprocessing.InputPath, imageprocessing.OutputSharpenPath)

	// Pass ProcessImageSharpenOptimized() to TimerWrapper()
	timedProcessImageSharpenOptimized := common.TimerWrapper(imageprocessing.ProcessImageSharpenOptimized)
	result4 := timedProcessImageSharpenOptimized(imageprocessing.InputPath, imageprocessing.OutputSharpenOptimizedPath)

	// Print results
	common.PrintResults(result1, result2, result3, result4)
}
