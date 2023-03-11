package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	bs := BubbleSort{}
	notSorted := []int{1, 3, 2}
	sorted := bs.Sort(notSorted)
	expectedSorted := []int{1, 2, 3}
	assert.Equal(t, expectedSorted, sorted)
}

func TestBubbleSortFloat(t *testing.T) {
	obj := new(BubbleSort)

	notSorted := []float64{5.5, 0.5, 1.2, 6.7, 0.3, 9.8}

	sorted := obj.SortFloat(notSorted)

	expected := []float64{0.3, 0.5, 1.2, 5.5, 6.7, 9.8}

	assert.Equal(t, expected, sorted)
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := BubbleSort{}
		bs.Sort([]int{1, 3, 2})
	}
}
