package median

import (
	"sort"
	"math"
	"errors"
)

func FindMedian(arr []int) (median float64, err error) {
	copied := make([]int, len(arr))
	copy(copied, arr)
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
