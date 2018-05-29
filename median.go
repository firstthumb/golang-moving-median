package main

import (
	"fmt"
	"sort"
	"math"
	"github.com/pkg/errors"
)

type SlidingWindow struct {
	data   []int
	size   int
	cursor int
}

func New(size int) *SlidingWindow {
	w := new(SlidingWindow)
	w.data = make([]int, size)
	w.size = size
	w.cursor = 0
	return w
}

func (w *SlidingWindow) Add(val int) {
	w.data[w.cursor] = val
	w.moveForward()
}

func (w *SlidingWindow) moveForward() {
	w.moveForwardBy(1)
}

func (w *SlidingWindow) moveForwardBy(amount int) {
	w.cursor = (w.cursor + amount) % w.size
}

func (w *SlidingWindow) FindMedian() (median float64, err error) {
	copied := make([]int, len(w.data))
	copy(copied, w.data)
	sort.Ints(copied)

	l := len(copied)
	if l == 0 {
		return math.NaN(), errors.New("No Data")
	} else if l%2 == 0 {
		median = float64(copied[l/2-1] + copied[l/2]) / 2
	} else {
		median = float64(copied[l/2])
	}

	return median, nil
}

func (w *SlidingWindow) Print() {
	fmt.Print("Array => ")
	for _, d := range w.data {
		fmt.Printf("%d ", d)
	}
	fmt.Println("")
}
