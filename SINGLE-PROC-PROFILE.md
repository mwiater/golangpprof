4 Goroutines:

```
Function                         |File Size       |Execution Time   |Performance Gain   |Concurrency
ProcessImageGrayscale            |5598865 Bytes   |772ms            |(baseline)         |1
ProcessImageGrayscaleOptimized   |5598865 Bytes   |520ms            |1.48x              |8
ProcessImageSharpen              |5598865 Bytes   |2590ms           |(baseline)         |1
ProcessImageSharpenOptimized     |5598865 Bytes   |838ms            |3.09x              |8

Max Baseline Throughput Per Day (GB):  268.01
Max Optimized Throughput Per Day (GB): 663.50
```

1 Goroutines:

```
Function                         |File Size       |Execution Time   |Performance Gain   |Concurrency
ProcessImageGrayscale            |5598865 Bytes   |788ms            |(baseline)         |1
ProcessImageGrayscaleOptimized   |5598865 Bytes   |812ms            |0.97x              |1
ProcessImageSharpen              |5598865 Bytes   |2583ms           |(baseline)         |1
ProcessImageSharpenOptimized     |5598865 Bytes   |2543ms           |1.02x              |1

Max Baseline Throughput Per Day (GB):  267.29
Max Optimized Throughput Per Day (GB): 268.57
```

`clear && go tool pprof ./pprof/cpu-ProcessImageGrayscaleOptimized.pprof`



list ProcessImageGrayscaleOptimized|toGrayscaleConcurrent

```
File: imageprocessing
Type: cpu
Time: Sep 24, 2023 at 7:04am (PDT)
Duration: 1.01s, Total samples = 790ms (78.00%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list ProcessImageGrayscaleOptimized|toGrayscaleConcurrent
Total: 790ms
ROUTINE ======================== imageprocessing.ProcessImageGrayscaleOptimized
         0      430ms (flat, cum) 54.43% of Total
         .          .     32:   return size, nil
         .          .     33:}
         .          .     34:
         .          .     35:func ProcessImageGrayscaleOptimized(inputPath string, outputPath string) (int64, error) {
         .          .     36:   size := getFileSize(inputPath)
         .      300ms     37:   img, err := decodeJPEG(inputPath)
         .          .     38:   if err != nil {
         .          .     39:           return 0, err
         .          .     40:   }
         .          .     41:
         .          .     42:   bounds := img.Bounds()
         .          .     43:   processedImage := image.NewGray(bounds)
         .          .     44:
         .          .     45:   var wg sync.WaitGroup
         .          .     46:   step := bounds.Dy() / NumRoutines
         .          .     47:
         .          .     48:   for i := 0; i < NumRoutines; i++ {
         .          .     49:           startY := i * step
         .          .     50:           endY := (i + 1) * step
         .          .     51:           if i == NumRoutines-1 {
         .          .     52:                   endY = bounds.Max.Y
         .          .     53:           }
         .          .     54:           wg.Add(1)
         .          .     55:           go toGrayscaleConcurrent(img, startY, endY, &wg, processedImage)
         .          .     56:   }
         .          .     57:   wg.Wait()
         .          .     58:
         .      130ms     59:   err = saveProcessedGrayScaleJPEG(outputPath, processedImage)
         .          .     60:   if err != nil {
         .          .     61:           return 0, err
         .          .     62:   }
         .          .     63:
         .          .     64:   return size, nil
ROUTINE ======================== imageprocessing.toGrayscaleConcurrent
      50ms      360ms (flat, cum) 45.57% of Total
         .          .     68:   defer wg.Done()
         .          .     69:
         .          .     70:   bounds := img.Bounds()
         .          .     71:   for y := start; y < end; y++ {
         .          .     72:           for x := bounds.Min.X; x < bounds.Max.X; x++ {
         .      160ms     73:                   c := img.At(x, y)
         .       30ms     74:                   gray, _, _, _ := c.RGBA()
      50ms      170ms     75:                   grayImage.Set(x, y, color.Gray{Y: uint8(gray >> 8)})
         .          .     76:           }
         .          .     77:   }
         .          .     78:}
```


```
File: imageprocessing
Type: cpu
Time: Sep 24, 2023 at 7:04am (PDT)
Duration: 1.01s, Total samples = 790ms (78.00%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) show=ProcessImageGrayscaleOptimized|toGrayscaleConcurrent
(pprof) top
Active filters:
   show=ProcessImageGrayscaleOptimized|toGrayscaleConcurrent
Showing nodes accounting for 790ms, 100% of 790ms total
      flat  flat%   sum%        cum   cum%
     430ms 54.43% 54.43%      430ms 54.43%  imageprocessing.ProcessImageGrayscaleOptimized
     360ms 45.57%   100%      360ms 45.57%  imageprocessing.toGrayscaleConcurrent
```


`clear && go tool pprof ./pprof/cpu-ProcessImageSharpenOptimized.pprof`


list ProcessImageSharpenOptimized|sharpenConcurrent

```
File: imageprocessing
Type: cpu
Time: Sep 24, 2023 at 7:04am (PDT)
Duration: 2.72s, Total samples = 2.55s (93.86%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list ProcessImageSharpenOptimized|sharpenConcurrent
Total: 2.55s
ROUTINE ======================== imageprocessing.ProcessImageSharpenOptimized
         0      570ms (flat, cum) 22.35% of Total
         .          .     50:   return size, nil
         .          .     51:}
         .          .     52:
         .          .     53:func ProcessImageSharpenOptimized(inputPath string, outputPath string) (int64, error) {
         .          .     54:   size := getFileSize(inputPath)
         .      290ms     55:   img, err := decodeJPEG(inputPath)
         .          .     56:   if err != nil {
         .          .     57:           return 0, err
         .          .     58:   }
         .          .     59:
         .          .     60:   bounds := img.Bounds()
         .       10ms     61:   processedImage := image.NewRGBA(bounds)
         .          .     62:
         .          .     63:   var wg sync.WaitGroup
         .          .     64:   step := bounds.Dy() / NumRoutines
         .          .     65:
         .          .     66:   for i := 0; i < NumRoutines; i++ {
         .          .     67:           startY := i * step
         .          .     68:           endY := (i + 1) * step
         .          .     69:           if i == NumRoutines-1 {
         .          .     70:                   endY = bounds.Max.Y
         .          .     71:           }
         .          .     72:           wg.Add(1)
         .          .     73:           go sharpenConcurrent(img, startY, endY, &wg, processedImage)
         .          .     74:   }
         .          .     75:   wg.Wait()
         .          .     76:
         .      270ms     77:   err = saveProcessedJPEG(outputPath, processedImage)
         .          .     78:   if err != nil {
         .          .     79:           return 0, err
         .          .     80:   }
         .          .     81:
         .          .     82:   return size, nil
ROUTINE ======================== imageprocessing.sharpenConcurrent
     460ms      1.95s (flat, cum) 76.47% of Total
         .          .     86:   defer wg.Done()
         .          .     87:   bounds := img.Bounds()
         .          .     88:   width := bounds.Dx()
         .          .     89:
         .          .     90:   for y := start; y < end; y++ {
      10ms       10ms     91:           for x := 1; x < width-1; x++ {
         .          .     92:                   var rSum, gSum, bSum int
         .          .     93:                   for ky := -1; ky <= 1; ky++ {
      40ms       40ms     94:                           for kx := -1; kx <= 1; kx++ {
      80ms      1.47s     95:                                   r, g, b, _ := img.At(x+kx, y+ky).RGBA()
     160ms      160ms     96:                                   kernelValue := sharpenKernel[ky+1][kx+1]
      70ms       70ms     97:                                   rSum += int(r) * kernelValue
      50ms       50ms     98:                                   gSum += int(g) * kernelValue
      20ms       20ms     99:                                   bSum += int(b) * kernelValue
         .          .    100:                           }
         .          .    101:                   }
         .          .    102:                   rValue := uint8(min(max(rSum>>8, 0), 255))
      10ms       10ms    103:                   gValue := uint8(min(max(gSum>>8, 0), 255))
      10ms       10ms    104:                   bValue := uint8(min(max(bSum>>8, 0), 255))
      10ms      110ms    105:                   output.Set(x, y, color.RGBA{R: rValue, G: gValue, B: bValue, A: 255})
         .          .    106:           }
         .          .    107:   }
         .          .    108:}
```

``````
File: imageprocessing
Type: cpu
Time: Sep 24, 2023 at 7:04am (PDT)
Duration: 2.72s, Total samples = 2.55s (93.86%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) show=ProcessImageSharpenOptimized|sharpenConcurrent
(pprof) top
Active filters:
   show=ProcessImageSharpenOptimized|sharpenConcurrent
Showing nodes accounting for 2520ms, 98.82% of 2550ms total
      flat  flat%   sum%        cum   cum%
    1950ms 76.47% 76.47%     1950ms 76.47%  imageprocessing.sharpenConcurrent
     570ms 22.35% 98.82%      570ms 22.35%  imageprocessing.ProcessImageSharpenOptimized
```