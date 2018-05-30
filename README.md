# Sliding Window Median

Sliding Window is circular array implementation. 
After the capacity is full, it starts to overwrite on existing value

```go
import (
    "movingmedian/slidingwindow"
) 

windowSize := 3
w := slidingwindow.New(windowSize)
w.add(100)
w.add(10)
w.add(20)
w.add(50)
w.add(40)
w.Print()
// [50, 40, 20]

```

Median package has simple function that takes array of int and returns median value

```go
import (
    "movingmedian/median"
) 

median, err := median.FindMedian([]int{10, 4, 6})
// median => 6

```

To run the application,

```bash
go run main.go
```

As you see in the benchmark test, If you increase window size, computational power drops dramatically.
Because the array size to handle is easily getting bigger.

```bash
~/g/s/movingmedian ❯❯❯ go test -bench=.                                                                                                             master ✱ ◼
goos: darwin
goarch: amd64
pkg: movingmedian
BenchmarkSlidingWindow_FindMedian_Size_2-8                    30          41804364 ns/op
BenchmarkSlidingWindow_FindMedian_Size_3-8                    30          43295491 ns/op
BenchmarkSlidingWindow_FindMedian_Size_4-8                    30          46423014 ns/op
BenchmarkSlidingWindow_FindMedian_Size_10-8                   20          66847744 ns/op
BenchmarkSlidingWindow_FindMedian_Size_20-8                   10         111908927 ns/op
BenchmarkSlidingWindow_FindMedian_Size_100-8                   2         624249145 ns/op
BenchmarkSlidingWindow_FindMedian_Size_150-8                   1        1019697768 ns/op
BenchmarkSlidingWindow_FindMedian_Size_500-8                   1        4270036362 ns/op
BenchmarkSlidingWindow_FindMedian_Size_1000-8                  1        9270674824 ns/op
PASS
ok      movingmedian    24.144s
```