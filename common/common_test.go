package common

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mwiater/golangpprof/imageprocessing"
	"github.com/stretchr/testify/assert"
)

const (
	testInput  = "../imageprocessing/inputs/input.jpg"
	testOutput = "../imageprocessing/outputs/grayscaleProcessed.jpg"
)

// Mock function for TimerWrapper test
func mockFunction(input, output string) (int64, error) {
	return 12345, nil
}

func TestGetFunctionName(t *testing.T) {
	assert.Equal(t, "mockFunction", getFunctionName(mockFunction))
}

func TestPrintResults(t *testing.T) {
	// Sample data for PrintResults function
	result1 := FunctionResult{
		FunctionName: "Function1",
		FileSize:     1000,
		Duration:     10,
	}

	result2 := FunctionResult{
		FunctionName: "Function2",
		FileSize:     2000,
		Duration:     5,
	}

	// Redirect standard output to capture the printed results
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintResults(result1, result2, FunctionResult{}, FunctionResult{})

	// Close the writer and restore standard output
	w.Close()
	os.Stdout = oldStdout

	// Read the captured output
	var buf bytes.Buffer
	buf.ReadFrom(r)

	// Check for specific outputs
	output := buf.String()

	assert.Contains(t, output, "Function1")
	assert.Contains(t, output, "1000 Bytes")
	assert.Contains(t, output, "10ms")
	assert.Contains(t, output, "(baseline)")
	assert.Contains(t, output, "1") // Concurrency for Function1

	assert.Contains(t, output, "Function2")
	assert.Contains(t, output, "2000 Bytes")
	assert.Contains(t, output, "5ms")
	assert.True(t, strings.Contains(output, "2.00x"))                   // Performance gain for Function2
	assert.Contains(t, output, fmt.Sprint(imageprocessing.NumRoutines)) // Concurrency for Function2
}
