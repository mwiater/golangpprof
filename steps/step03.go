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

	// Not running these yet, passing placeholders
	result3 := common.FunctionResult{}
	result4 := common.FunctionResult{}

	// Print results
	common.PrintResults(result1, result2, result3, result4)
}
