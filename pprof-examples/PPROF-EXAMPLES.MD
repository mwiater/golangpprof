# PPROF-EXAMPLES

**Reference data for Article: [Pprof Through Examples: Exploring Optimizations in Go](https://medium.com/@matt.wiater/pprof-through-examples-exploring-optimizations-in-go-444fa08cf15f)**

## Output

```
Function                       |File Size     |Execution Time   |Perf. Gain   |Concurrency
ProcessImageGrayscale          |5598865 Bytes |777ms            |(baseline)   |1
ProcessImageGrayscaleOptimized |5598865 Bytes |501ms            |1.55x        |8
ProcessImageSharpen            |5598865 Bytes |2624ms           |(baseline)   |1
ProcessImageSharpenOptimized   |5598865 Bytes |876ms            |3.00x        |8

Max Baseline Throughput Per Day (GB):  264.93
Max Optimized Throughput Per Day (GB): 654.35
```

## Pprofs
`go tool pprof ./pprof-examples/cpu-ProcessImageGrayscale.pprof`

`list ProcessImageGrayscale`

```
File: imageprocessing
Type: cpu
Time: Sep 30, 2023 at 10:27am (PDT)
Duration: 914.10ms, Total samples = 760ms (83.14%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list ProcessImageGrayscale
Total: 760ms
ROUTINE ======================== imageprocessing/ProcessImageGrayscale
    0      760ms (flat, cum)   100% of Total
    .          .     29:func ProcessImageGrayscale(inputPath string, outputPath string) (int64, error) {
    .          .     30:   size, err := getFileSize(inputPath)
    .          .     31:   if err != nil {
    .          .     32:     return 0, err
    .          .     33:   }
    .      290ms     34:   img, err := decodeJPEG(inputPath)
    .          .     35:   if err != nil {
    .          .     36:     return 0, err
    .          .     37:   }
    .          .     38:
    .          .     39:   bounds := img.Bounds()
    .          .     40:   processedImage := image.NewGray(bounds)
    .          .     41:
    .          .     42:   for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    .          .     43:     for x := bounds.Min.X; x < bounds.Max.X; x++ {
    .      180ms     44:       c := img.At(x, y)
    .          .     45:       gray, _, _, _ := c.RGBA()
    .      150ms     46:       processedImage.Set(x, y, color.Gray{Y: uint8(gray >> 8)})
    .          .     47:     }
    .          .     48:   }
    .          .     49:
    .      140ms     50:   err = saveProcessedGrayScaleJPEG(outputPath, processedImage)
    .          .     51:   if err != nil {
    .          .     52:     return 0, err
    .          .     53:   }
    .          .     54:
    .          .     55:   return size, nil
    .          .     56:}
```

-------------------


`go tool pprof ./pprof-examples/cpu-ProcessImageGrayscaleOptimized.pprof`

`list ProcessImageGrayscaleOptimized|toGrayscaleConcurrent`

```
File: imageprocessing
Type: cpu
Time: Sep 30, 2023 at 10:27am (PDT)
Duration: 616.89ms, Total samples = 790ms (128.06%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list ProcessImageGrayscaleOptimized|toGrayscaleConcurrent
Total: 790ms
ROUTINE ======================== imageprocessing/ProcessImageGrayscaleOptimized
    0      420ms (flat, cum) 53.16% of Total
    .          .     81:func ProcessImageGrayscaleOptimized(inputPath string, outputPath string) (int64, error) {
    .          .     82:   size, err := getFileSize(inputPath)
    .          .     83:   if err != nil {
    .          .     84:     return 0, err
    .          .     85:   }
    .      290ms     86:   img, err := decodeJPEG(inputPath)
    .          .     87:   if err != nil {
    .          .     88:     return 0, err
    .          .     89:   }
    .          .     90:
    .          .     91:   bounds := img.Bounds()
    .          .     92:   processedImage := image.NewGray(bounds)
    .          .     93:
    .          .     94:   var wg sync.WaitGroup
    .          .     95:   step := bounds.Dy() / NumRoutines
    .          .     96:
    .          .     97:   for i := 0; i < NumRoutines; i++ {
    .          .     98:     startY := i * step
    .          .     99:     endY := (i + 1) * step
    .          .    100:     if i == NumRoutines-1 {
    .          .    101:       endY = bounds.Max.Y
    .          .    102:     }
    .          .    103:     wg.Add(1)
    .          .    104:     go toGrayscaleConcurrent(img, startY, endY, &wg, processedImage)
    .          .    105:   }
    .          .    106:   wg.Wait()
    .          .    107:
    .      130ms    108:   err = saveProcessedGrayScaleJPEG(outputPath, processedImage)
    .          .    109:   if err != nil {
    .          .    110:     return 0, err
    .          .    111:   }
    .          .    112:
    .          .    113:   return size, nil
    .          .    114:}
ROUTINE ======================== imageprocessing/toGrayscaleConcurrent
 30ms      370ms (flat, cum) 46.84% of Total
    .          .    128:func toGrayscaleConcurrent(img image.Image, start, end int, wg *sync.WaitGroup, grayImage *image.Gray) {
    .          .    129:   defer wg.Done()
    .          .    130:
    .          .    131:   bounds := img.Bounds()
    .          .    132:   for y := start; y < end; y++ {
    .          .    133:     for x := bounds.Min.X; x < bounds.Max.X; x++ {
 10ms      170ms    134:       c := img.At(x, y)
 10ms       40ms    135:       gray, _, _, _ := c.RGBA()
 10ms      160ms    136:       grayImage.Set(x, y, color.Gray{Y: uint8(gray >> 8)})
    .          .    137:     }
    .          .    138:   }
    .          .    139:}
```

-------------------

`go tool pprof ./pprof-examples/cpu-ProcessImageSharpen.pprof`

`list ProcessImageSharpen`

```
File: imageprocessing
Type: cpu
Time: Sep 30, 2023 at 10:27am (PDT)
Duration: 2.82s, Total samples = 2.60s (92.32%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list ProcessImageSharpen
Total: 2.60s
ROUTINE ======================== imageprocessing/ProcessImageSharpen
460ms      2.60s (flat, cum)   100% of Total
    .          .     38:func ProcessImageSharpen(inputPath string, outputPath string) (int64, error) {
    .          .     39:   size, err := getFileSize(inputPath)
    .          .     40:   if err != nil {
    .          .     41:     return 0, err
    .          .     42:   }
    .      310ms     43:   img, err := decodeJPEG(inputPath)
    .          .     44:   if err != nil {
    .          .     45:     return 0, err
    .          .     46:   }
    .          .     47:
    .          .     48:   bounds := img.Bounds()
    .          .     49:   width, height := bounds.Dx(), bounds.Dy()
    .          .     50:   processedImage := image.NewRGBA(bounds)
    .          .     51:
    .          .     52:   for y := 1; y < height-1; y++ {
    .          .     53:     for x := 1; x < width-1; x++ {
    .          .     54:       var rSum, gSum, bSum int
 10ms       10ms     55:       for ky := -1; ky <= 1; ky++ {
 20ms       20ms     56:         for kx := -1; kx <= 1; kx++ {
 60ms      1.46s     57:           r, g, b, _ := img.At(x+kx, y+ky).RGBA()
240ms      240ms     58:           kernelValue := sharpenKernel[ky+1][kx+1]
 40ms       40ms     59:           rSum += int(r) * kernelValue
 10ms       10ms     60:           gSum += int(g) * kernelValue
 10ms       20ms     61:           bSum += int(b) * kernelValue
    .          .     62:         }
    .          .     63:       }
    .          .     64:       rValue := uint8(min(max(rSum>>8, 0), 255))
    .          .     65:       gValue := uint8(min(max(gSum>>8, 0), 255))
 30ms       30ms     66:       bValue := uint8(min(max(bSum>>8, 0), 255))
 40ms      180ms     67:       processedImage.Set(x, y, color.RGBA{R: rValue, G: gValue, B: bValue, A: 255})
    .          .     68:     }
    .          .     69:   }
    .          .     70:
    .      280ms     71:   err = saveProcessedJPEG(outputPath, processedImage)
    .          .     72:   if err != nil {
    .          .     73:     return 0, err
    .          .     74:   }
    .          .     75:
    .          .     76:   return size, nil
    .          .     77:}
```

-------------------

`go tool pprof ./pprof-examples/cpu-ProcessImageSharpenOptimized.pprof`

`list ProcessImageSharpenOptimized|sharpenConcurrent`

```
File: imageprocessing
Type: cpu
Time: Sep 30, 2023 at 10:27am (PDT)
Duration: 1.03s, Total samples = 2.73s (264.91%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list ProcessImageSharpenOptimized|sharpenConcurrent
Total: 2.73s
ROUTINE ======================== imageprocessing/ProcessImageSharpenOptimized
    0      580ms (flat, cum) 21.25% of Total
    .          .    103:func ProcessImageSharpenOptimized(inputPath string, outputPath string) (int64, error) {
    .          .    104:   size, err := getFileSize(inputPath)
    .          .    105:   if err != nil {
    .          .    106:     return 0, err
    .          .    107:   }
    .      280ms    108:   img, err := decodeJPEG(inputPath)
    .          .    109:   if err != nil {
    .          .    110:     return 0, err
    .          .    111:   }
    .          .    112:
    .          .    113:   bounds := img.Bounds()
    .          .    114:   processedImage := image.NewRGBA(bounds)
    .          .    115:
    .          .    116:   var wg sync.WaitGroup
    .          .    117:   step := bounds.Dy() / NumRoutines
    .          .    118:
    .          .    119:   for i := 0; i < NumRoutines; i++ {
    .          .    120:     startY := i * step
    .          .    121:     endY := (i + 1) * step
    .          .    122:     if i == NumRoutines-1 {
    .          .    123:       endY = bounds.Max.Y
    .          .    124:     }
    .          .    125:     wg.Add(1)
    .          .    126:     go sharpenConcurrent(img, startY, endY, &wg, processedImage)
    .          .    127:   }
    .          .    128:   wg.Wait()
    .          .    129:
    .      300ms    130:   err = saveProcessedJPEG(outputPath, processedImage)
    .          .    131:   if err != nil {
    .          .    132:     return 0, err
    .          .    133:   }
    .          .    134:
    .          .    135:   return size, nil
    .          .    136:}
ROUTINE ======================== imageprocessing/sharpenConcurrent
460ms      2.14s (flat, cum) 78.39% of Total
    .          .    155:func sharpenConcurrent(img image.Image, start, end int, wg *sync.WaitGroup, output *image.RGBA) {     
    .          .    156:   defer wg.Done()
    .          .    157:   bounds := img.Bounds()
    .          .    158:   width := bounds.Dx()
    .          .    159:
    .          .    160:   for y := start; y < end; y++ {
 10ms       10ms    161:     for x := 1; x < width-1; x++ {
    .          .    162:       var rSum, gSum, bSum int
 10ms       10ms    163:       for ky := -1; ky <= 1; ky++ {
 10ms       10ms    164:         for kx := -1; kx <= 1; kx++ {
160ms      1.73s    165:           r, g, b, _ := img.At(x+kx, y+ky).RGBA()
150ms      150ms    166:           kernelValue := sharpenKernel[ky+1][kx+1]
 30ms       30ms    167:           rSum += int(r) * kernelValue
    .          .    168:           gSum += int(g) * kernelValue
 30ms       30ms    169:           bSum += int(b) * kernelValue
    .          .    170:         }
    .          .    171:       }
    .       10ms    172:       rValue := uint8(min(max(rSum>>8, 0), 255))
    .          .    173:       gValue := uint8(min(max(gSum>>8, 0), 255))
 30ms       30ms    174:       bValue := uint8(min(max(bSum>>8, 0), 255))
 30ms      130ms    175:       output.Set(x, y, color.RGBA{R: rValue, G: gValue, B: bValue, A: 255})
    .          .    176:     }
    .          .    177:   }
    .          .    178:}
```
