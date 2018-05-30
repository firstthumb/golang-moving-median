package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"movingmedian/slidingwindow"
	"movingmedian/median"
)


func TestSlidingWindow_FindMedian(t *testing.T) {
	w := slidingwindow.New(3)
	w.Add(3)
	w.Add(1)
	w.Add(5)
	w.Add(7)
	w.Add(3)
	w.Add(2)
	w.Add(6)
	w.Add(4)

	medianValue, err := median.FindMedian(w.Data)
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, medianValue, float64(4), "Median is wrong")
}

func TestSlidingWindow_FindMedian2(t *testing.T) {
	w := slidingwindow.New(4)
	w.Add(3)
	w.Add(1)
	w.Add(5)
	w.Add(7)
	w.Add(3)
	w.Add(2)
	w.Add(6)
	w.Add(4)

	medianValue, err := median.FindMedian(w.Data)
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, medianValue, float64(3.5), "Median is wrong")
}
