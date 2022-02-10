package formatters

import (
	"testing"

	"github.com/linden/chroma"
	"github.com/stretchr/testify/assert"
)

func TestClosestColour(t *testing.T) {
	actual := findClosest(ttyTables[256], chroma.MustParseColour("#e06c75"))
	assert.Equal(t, chroma.MustParseColour("#d75f87"), actual)
}
