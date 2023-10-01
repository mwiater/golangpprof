package main

import (
	"github.com/mwiater/golangpprof/common"
	"github.com/mwiater/golangpprof/imageprocessing"
)

func main() {
	// Pass ProcessImageGrayscale() to TimerWrapper
	timedProcessImageGrayscale := common.TimerWrapper(imageprocessing.ProcessImageGrayscale)
	result1 := timedProcessImageGrayscale(imageprocessing.InputPath, imageprocessing.OutputGrayscalePath)

	// Not running these yet, passing placeholders
	result2 := common.FunctionResult{}
	result3 := common.FunctionResult{}
	result4 := common.FunctionResult{}

	// Print Results
	common.PrintResults(result1, result2, result3, result4)
}
