package imageprocessing

import "runtime"

const (
	InputPath                    = "./imageprocessing/inputs/input.jpg"
	OutputGrayscalePath          = "./imageprocessing/outputs/grayscaleProcessed.jpg"
	OutputGrayscaleOptimizedPath = "./imageprocessing/outputs/grayscaleProcessedOptimized.jpg"
	OutputSharpenPath            = "./imageprocessing/outputs/sharpenProcessed.jpg"
	OutputSharpenOptimizedPath   = "./imageprocessing/outputs/sharpenProcessedOptimized.jpg"
)

var NumRoutines = runtime.NumCPU()
