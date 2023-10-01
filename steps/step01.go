package main

import (
	"fmt"

	"github.com/mwiater/golangpprof/imageprocessing"
)

func main() {
	_, err := imageprocessing.ProcessImageGrayscale(imageprocessing.InputPath, imageprocessing.OutputGrayscalePath)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}
}
