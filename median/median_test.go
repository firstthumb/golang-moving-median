package median

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlidingWindow_FindMedian(t *testing.T) {
	arr := []int{2, 6, 4}

	median, err := FindMedian(arr)
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, median, float64(4), "Median is wrong")
}

func TestSlidingWindow_FindMedian2(t *testing.T) {
	arr := []int{3, 2, 6, 4}

	median, err := FindMedian(arr)
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, median, float64(3.5), "Median is wrong")
}

