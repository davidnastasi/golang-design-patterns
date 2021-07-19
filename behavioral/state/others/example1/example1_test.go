package example1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	conn := &Connection{CloseState{}}
	assert.NoError(t, conn.Open())
	assert.Error(t, conn.Open())
	assert.NoError(t, conn.Close())
	assert.Error(t, conn.Close())

}
