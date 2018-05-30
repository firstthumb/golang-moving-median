package slidingwindow

import "fmt"

type SlidingWindow struct {
	data   	[]int
	size   	int
	cursor 	int
	length	int
}

func New(size int) *SlidingWindow {
	w := new(SlidingWindow)
	w.data = make([]int, size)
	w.size = size
	w.length = 0
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
	w.length = w.length + amount
}

func (w *SlidingWindow) Length() int {
	if w.length > w.size {
		return w.size
	}

	return w.length
}

func (w *SlidingWindow) Slice() []int {
	return w.data[:w.Length()]
}

func (w *SlidingWindow) Print() {
	fmt.Print("Array => ")
	for _, d := range w.data {
		fmt.Printf("%d ", d)
	}
	fmt.Println("")
}
