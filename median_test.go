package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow_Add(t *testing.T) {
	w := New(3)
	w.Add(1)
	assert.Equal(t, w.cursor, 1, "Expected only one item")

	w.Add(2)
	assert.Equal(t, w.cursor, 2, "Expected only two item")

	w.Add(3)
	assert.Equal(t, w.cursor, 0, "Expected only three item")

	w.Add(4)
	assert.Equal(t, w.cursor, 1, "Expected only three item")
	assert.Equal(t, len(w.data), 3, "Array should have 3 items")
	assert.Equal(t, w.data, []int{4, 2, 3}, "End data is wrong")

	w.Add(5)
	assert.Equal(t, w.data, []int{4, 5, 3}, "End data is wrong")
}

func TestSlidingWindow_FindMedian(t *testing.T) {
	w := New(3)
	w.Add(3)
	w.Add(1)
	w.Add(5)
	w.Add(7)
	w.Add(3)
	w.Add(2)
	w.Add(6)
	w.Add(4)

	median, err := w.FindMedian()
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, median, float64(4), "Median is wrong")
}

func TestSlidingWindow_FindMedian2(t *testing.T) {
	w := New(4)
	w.Add(3)
	w.Add(1)
	w.Add(5)
	w.Add(7)
	w.Add(3)
	w.Add(2)
	w.Add(6)
	w.Add(4)

	median, err := w.FindMedian()
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, median, float64(3.5), "Median is wrong")
}
