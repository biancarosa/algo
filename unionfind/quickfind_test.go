package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFind(t *testing.T) {
	quickfind := QuickFind{}
	quickfind.New(3)

	assert.False(t, quickfind.Connected(0, 1))
	assert.False(t, quickfind.Connected(1, 2))
	assert.False(t, quickfind.Connected(0, 2))

	quickfind.Union(1, 2)

	assert.False(t, quickfind.Connected(0, 1))
	assert.True(t, quickfind.Connected(1, 2))
	assert.False(t, quickfind.Connected(0, 2))

	quickfind.Union(0, 2)
	assert.True(t, quickfind.Connected(0, 1))
	assert.True(t, quickfind.Connected(1, 2))
	assert.True(t, quickfind.Connected(0, 2))
}

func BenchmarkQuickFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quickfind := QuickFind{}
		quickfind.New(3)
		quickfind.Union(1, 2)
		quickfind.Union(0, 2)
		quickfind.Connected(0, 1)
		quickfind.Connected(1, 2)
		quickfind.Connected(0, 2)
	}
}
