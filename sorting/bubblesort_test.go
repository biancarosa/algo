package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	bs := BubbleSort{}
	sorted := bs.Sort([]int{1, 3, 2})
	assert.Equal(t, []int{1, 2, 3}, sorted)
}
