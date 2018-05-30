package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"movingmedian/slidingwindow"
	"movingmedian/median"
	"math/rand"
	"time"
	"strconv"
)

func TestSlidingWindow_Main(t *testing.T) {
	w := slidingwindow.New(3)
	w.Add(3)
	w.Add(1)
	w.Add(5)
	w.Add(7)
	w.Add(3)
	w.Add(2)
	w.Add(6)
	w.Add(4)

	medianValue, err := median.FindMedian(w.Slice())
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, medianValue, float64(4), "Median is wrong")
}

func TestSlidingWindow_Main2(t *testing.T) {
	w := slidingwindow.New(4)
	w.Add(3)
	w.Add(1)
	w.Add(5)
	w.Add(7)
	w.Add(3)
	w.Add(2)
	w.Add(6)
	w.Add(4)

	medianValue, err := median.FindMedian(w.Slice())
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, medianValue, float64(3.5), "Median is wrong")
}

func benchmarkSlidingWindow_FindMedian(windowSize int, CalculateMedians func(int, []string) []string, b *testing.B) {
	inputs := make([]string, 100000)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(inputs); i++ {
		inputs[i] = strconv.Itoa(rand.Intn(200))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalculateMedians(windowSize, inputs)
	}
}

func BenchmarkSlidingWindow_FindMedian_Size_2(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(2, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_3(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(3, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_4(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(4, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_10(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(10, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_20(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(20, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_100(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(100, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_150(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(150, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_500(b *testing.B)      	{ benchmarkSlidingWindow_FindMedian(500, calculateMedians, b) }
func BenchmarkSlidingWindow_FindMedian_Size_1000(b *testing.B)      { benchmarkSlidingWindow_FindMedian(1000, calculateMedians, b) }