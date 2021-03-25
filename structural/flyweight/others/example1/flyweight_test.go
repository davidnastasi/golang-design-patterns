package example1

import (
	"github.com/stretchr/testify/require"
	"image/color"
	"testing"
)

func TestFlyweight(t *testing.T) {

	ct1 := NewCounterTerroristPlayer()
	ct2 := NewCounterTerroristPlayer()
	ct3 := NewCounterTerroristPlayer()

	t1 := NewTerroristPlayer()
	t2 := NewTerroristPlayer()
	t3 := NewTerroristPlayer()

	require.Equal(t, ct1.dresser.GetDress().Color, color.White)
	require.Equal(t, &ct1.dresser, &ct2.dresser)
	require.Equal(t, &ct2.dresser, &ct3.dresser)

	require.Equal(t, t1.dresser.GetDress().Color, color.Black)
	require.Equal(t, &t1.dresser, &t2.dresser)
	require.Equal(t, &t2.dresser, &t3.dresser)

}