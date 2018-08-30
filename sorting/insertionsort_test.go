package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	bs := InsertionSort{}
	sorted := bs.Sort([]int{1, 3, 2})
	assert.Equal(t, []int{1, 2, 3}, sorted)
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bs := InsertionSort{}
		bs.Sort([]int{1, 3, 2})
	}
}
