package slidingwindow

import "fmt"

type SlidingWindow struct {
	Data   []int
	size   int
	cursor int
}

func New(size int) *SlidingWindow {
	w := new(SlidingWindow)
	w.Data = make([]int, size)
	w.size = size
	w.cursor = 0
	return w
}

func (w *SlidingWindow) Add(val int) {
	w.Data[w.cursor] = val
	w.moveForward()
}

func (w *SlidingWindow) moveForward() {
	w.moveForwardBy(1)
}

func (w *SlidingWindow) moveForwardBy(amount int) {
	w.cursor = (w.cursor + amount) % w.size
}

func (w *SlidingWindow) Print() {
	fmt.Print("Array => ")
	for _, d := range w.Data {
		fmt.Printf("%d ", d)
	}
	fmt.Println("")
}
