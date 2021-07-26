package example1

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGame_Rook(t *testing.T) {
	rook := &Rook{}
	paper := &Paper{}
	scissor := &Scissor{}
	rookVisitor := &RookVisitor{}
	require.Equal(t, DRAW, rook.Accept(rookVisitor))
	require.Equal(t, LOOSE, paper.Accept(rookVisitor))
	require.Equal(t, WIN, scissor.Accept(rookVisitor))
}

