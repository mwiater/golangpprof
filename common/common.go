package common

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"runtime/pprof"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mwiater/golangpprof/imageprocessing"
)

// FunctionResult is a structure representing the result from an image processing function.
type FunctionResult struct {
	FunctionName string
	FileSize     int64
	Duration     float64
}

// WrappedImageProcessingFunction is a function signature for image processing functions that can be wrapped by the TimerWrapper.
type WrappedImageProcessingFunction func(string, string) (int64, error)

// TimerWrapper takes an image processing function and wraps it to measure and report its execution time and CPU profiling.
//
// Parameters:
// - fn: The image processing function to be wrapped.
//
// Returns:
// - A new function with the same signature as the input function, but returns a FunctionResult instead of the usual (int64, error).
func TimerWrapper(fn WrappedImageProcessingFunction) func(string, string) FunctionResult {
	return func(inputPath string, outputPath string) FunctionResult {
		functionName := getFunctionName(fn)
		fmt.Println("Profiling: " + functionName + "()")
		cpuprofile1 := "./pprof/cpu-" + functionName + ".pprof"
		f, err := os.Create(cpuprofile1)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not create CPU profile: %v\n", err)
			panic(err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "could not start CPU profile: %v\n", err)
			panic(err)
		}
		start := time.Now()
		fileSize, err := fn(inputPath, outputPath)
		if err != nil {
			panic(err)
		}
		elapsed := time.Since(start)
		duration := float64(elapsed.Milliseconds())

		pprof.StopCPUProfile()
		fmt.Println("  ...Complete")
		fmt.Println()

		return FunctionResult{
			FunctionName: functionName,
			FileSize:     fileSize,
			Duration:     duration,
		}
	}
}

// getFunctionName retrieves the name of the provided function.
//
// Parameters:
// - i: The interface whose underlying function name is to be retrieved.
//
// Returns:
// - string: The name of the function.
func getFunctionName(i interface{}) string {
	ptr := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return strings.Split(ptr, ".")[len(strings.Split(ptr, "."))-1]
}

// PrintResults prints out the results of the image processing functions in a tabulated format.
//
// Parameters:
// - result1, result2, result3, result4: FunctionResults to be printed.
func PrintResults(result1 FunctionResult, result2 FunctionResult, result3 FunctionResult, result4 FunctionResult) {
	fmt.Println()
	w := tabwriter.NewWriter(os.Stdout, 10, 1, 3, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", "Function", "File Size", "Execution Time", "Performance Gain", "Concurrency")
	fmt.Fprintf(w, "%s\t%d Bytes\t%.0fms\t%s\t%d\n", result1.FunctionName, result1.FileSize, result1.Duration, "(baseline)", 1)
	if (result2 != FunctionResult{}) {
		fmt.Fprintf(w, "%s\t%d Bytes\t%.0fms\t%.2fx\t%d\n", result2.FunctionName, result2.FileSize, result2.Duration, result1.Duration/result2.Duration, imageprocessing.NumRoutines)
	}
	if (result3 != FunctionResult{}) {
		fmt.Fprintf(w, "%s\t%d Bytes\t%.0fms\t%s\t%d\n", result3.FunctionName, result3.FileSize, result3.Duration, "(baseline)", 1)
	}
	if (result4 != FunctionResult{}) {
		fmt.Fprintf(w, "%s\t%d Bytes\t%.0fms\t%.2fx\t%d\n", result4.FunctionName, result4.FileSize, result4.Duration, result3.Duration/result4.Duration, imageprocessing.NumRoutines)
	}
	w.Flush()

	if (result2 != FunctionResult{} && result3 != FunctionResult{} && result4 != FunctionResult{}) {
		fmt.Println()
		totalBytesProcessedBaseline := result1.FileSize + result3.FileSize
		totalBytesOptimizedOptimized := result2.FileSize + result4.FileSize
		totalTimeBaseline := result1.Duration + result3.Duration
		totalTimeOptimized := result2.Duration + result4.Duration

		bytesPerMillisecondBaseline := float64(totalBytesProcessedBaseline) / totalTimeBaseline
		bytesPerMillisecondOptimized := float64(totalBytesOptimizedOptimized) / totalTimeOptimized

		fmt.Printf("Max Baseline Throughput Per Day (GB):  %.2f\n", (bytesPerMillisecondBaseline*86400000)/1073741824)
		fmt.Printf("Max Optimized Throughput Per Day (GB): %.2f\n", (bytesPerMillisecondOptimized*86400000)/1073741824)
	}
	fmt.Println()
}
