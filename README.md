# golangpprof

**Source code and examples for the article: [Pprof Through Examples: Exploring Optimizations in Go](https://medium.com/@matt.wiater/pprof-through-examples-exploring-optimizations-in-go-444fa08cf15f)**

**Table of Contents**

* [Setup](#setup)
* [Build binaries](#build-binaries)
* [Run binaries](#run-binaries)
* [Tests](#tests)
* [Pprof Examples](#pprof-examples)

---

## Setup

```
git clone https://github.com/mwiater/golangpprof
cd golangpprof
go mod tidy
```

---

## Build binaries

First, build all of the binaries for a more direct, consistent run:

```
go build -o ./bin imageprocessing.go
go build -o ./bin steps/step01.go
go build -o ./bin steps/step02.go
go build -o ./bin steps/step03.go
go build -o ./bin steps/step04.go
```

---

## Run binaries

### imageprocessing

`./bin/imageprocessing`

Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`
Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`
Runs: ProcessImageSharpen
Outputs: `./pprof/cpu-ProcessImageSharpen.pprof`
Runs: ProcessImageSharpenOptimized
Outputs: `./pprof/cpu-ProcessImageSharpenOptimized.pprof`

Example command output:

```
Function                         |File Size       |Execution Time   |Performance Gain   |Concurrency
ProcessImageGrayscale            |5598865 Bytes   |772ms            |(baseline)         |1
ProcessImageGrayscaleOptimized   |5598865 Bytes   |520ms            |1.48x              |8
ProcessImageSharpen              |5598865 Bytes   |2590ms           |(baseline)         |1
ProcessImageSharpenOptimized     |5598865 Bytes   |838ms            |3.09x              |8

Max Baseline Throughput Per Day (GB):  268.01
Max Optimized Throughput Per Day (GB): 663.50
```

### step01

`./bin/step01`

Runs: ProcessImageGrayscale

Example command output:

```
Success
```

### step02

`./bin/step02`

Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`

Example command output:

```
Function                         |File Size       |Execution Time   |Performance Gain   |Concurrency
ProcessImageGrayscale            |5598865 Bytes   |772ms            |(baseline)         |1
```

### step03

`./bin/step03`

Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`
Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`

Example command output:

```
Function                         |File Size       |Execution Time   |Performance Gain   |Concurrency
ProcessImageGrayscale            |5598865 Bytes   |772ms            |(baseline)         |1
ProcessImageGrayscaleOptimized   |5598865 Bytes   |520ms            |1.48x              |8
```

### step04

`./bin/step04`

Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`
Runs: ProcessImageGrayscale
Outputs: `./pprof/cpu-ProcessImageGrayscale.pprof`
Runs: ProcessImageSharpen
Outputs: `./pprof/cpu-ProcessImageSharpen.pprof`
Runs: ProcessImageSharpenOptimized
Outputs: `./pprof/cpu-ProcessImageSharpenOptimized.pprof`

Example command output:

```
Function                         |File Size       |Execution Time   |Performance Gain   |Concurrency
ProcessImageGrayscale            |5598865 Bytes   |772ms            |(baseline)         |1
ProcessImageGrayscaleOptimized   |5598865 Bytes   |520ms            |1.48x              |8
ProcessImageSharpen              |5598865 Bytes   |2590ms           |(baseline)         |1
ProcessImageSharpenOptimized     |5598865 Bytes   |838ms            |3.09x              |8

Max Baseline Throughput Per Day (GB):  268.01
Max Optimized Throughput Per Day (GB): 663.50
```

---

## Tests

`go install gotest.tools/gotestsum@latest`

`go clean -testcache && gotestsum -f testname ./imageprocessing ./common`

```
PASS common.TestGetFunctionName (0.00s)
PASS common.TestPrintResults (0.00s)
PASS common
PASS imageprocessing.TestProcessImageSharpen_Decode (0.30s)
PASS imageprocessing.TestProcessImageSharpen_Processing (2.59s)
PASS imageprocessing.TestProcessImageSharpen_OutputSize (0.00s)
PASS imageprocessing.TestProcessImageSharpen_DifferentFromInput (0.89s)
PASS imageprocessing.TestProcessImageSharpenOptimized_Decode (0.30s)
PASS imageprocessing.TestProcessImageSharpenOptimized_Processing (0.90s)
PASS imageprocessing.TestProcessImageSharpenOptimized_OutputSize (0.00s)
PASS imageprocessing.TestProcessImageSharpenOptimized_DifferentFromInput (0.92s)
PASS imageprocessing.TestProcessImageGrayscale_Decode (0.78s)
PASS imageprocessing.TestProcessImageGrayscale_OutputSize (0.00s)
PASS imageprocessing.TestProcessImageGrayscale_GrayscaleOutput (0.19s)
PASS imageprocessing.TestProcessImageGrayscaleOptimized_Decode (0.51s)
PASS imageprocessing.TestProcessImageGrayscaleOptimized_OutputSize (0.00s)
PASS imageprocessing.TestProcessImageGrayscaleOptimized_GrayscaleOutput (0.19s)
PASS imageprocessing

DONE 16 tests in 7.836s
```

---

## Pprof Examples

Files included:

```
./pprof-examples/
  pprof-examples/cpu-ProcessImageGrayscale.pprof
  pprof-examples/cpu-ProcessImageGrayscaleOptimized.pprof
  pprof-examples/cpu-ProcessImageSharpen.pprof
  pprof-examples/cpu-ProcessImageSharpenOptimized.pprof
  pprof-examples/PPROF-EXAMPLES.MD
```

Since running the application generates slightly different results each time based on your systems current resource usage at the time of running, I've included example pprof files from my machine during a single run. This way our comparisons are not altered by variations that may be introduced during multiple runs. Variations like this are expected, but when illustrating with examples, eliminating these variations makes comparisons clearer.

**These are the files and output referenced in the article: [Pprof Through Examples: Exploring Optimizations in Go](https://medium.com/@matt.wiater/pprof-through-examples-exploring-optimizations-in-go-444fa08cf15f)**

See [pprof-examples/PPROF-EXAMPLES.MD](pprof-examples/PPROF-EXAMPLES.MD) for full output results.
