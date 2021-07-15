package example1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChristmasTree(t *testing.T) {
	christmasTreeDecorated := Decorate( &ChristmasTree{}, DecorateTreeTrooper(), DecorateTinsel(), DecorateGarland(), DecorateBubbleLights())
	assert.Equal(t, "ChristmasTree: TreeTrooper Tinsel Garland BubbleLights" , christmasTreeDecorated.Do())
}