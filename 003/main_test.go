package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlope(t *testing.T) {
	originalMap := `..##..
......
######
      `

	t.Run("get bigger map", func(t *testing.T) {

		expectedMap := `..##....##....##..
..................
##################
                 `

		got := getBiggerMap(originalMap, 3)

		require.Equal(t, expectedMap, got)
	})

	t.Run("count trees", func(t *testing.T) {
		iMap := getBiggerMap(originalMap, 3)

		require.Equal(t, 1, countTheTrees(iMap))
	})
}
