package median

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/**

If only one element available in the sliding window the answer is -1.
If n is odd then Median (M) = value of ((n + 1)/2)th item from a sorted array of length n.
If n is even then Median (M) = value of [((n)/2)th item term + ((n)/2 + 1)th item term ] /2

 */

func TestFindMedian(t *testing.T) {
	arr := []int{2, 6, 4}

	median, err := FindMedian(arr)
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, median, float64(4), "Median is wrong")
}

func TestFindMedian2(t *testing.T) {
	arr := []int{3, 2, 6, 4}

	median, err := FindMedian(arr)
	assert.Nil(t, err, "Error has been thrown")
	assert.Equal(t, median, float64(3.5), "Median is wrong")
}

func TestFindMedian3(t *testing.T) {
	arr := []int{}

	_, err := FindMedian(arr)
	assert.NotNil(t, err, "It should return error")
}

func TestFindMedian4(t *testing.T) {
	arr := []int{1}

	_, err := FindMedian(arr)
	assert.NotNil(t, err, "It should return error")
}

