package unionfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFind(t *testing.T) {
	quickfind := QuickFindImpl{}
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
