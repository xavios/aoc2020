package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	cases := map[string]struct {
		account    []int
		numOfparts int
		want       int
	}{
		"basic example": {
			[]int{1, 2019},
			2,
			2019,
		},
		"basic example of three": {
			[]int{1, 2018, 1},
			3,
			2018,
		},
		"provided example": {
			[]int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			2,
			514579,
		},
		"provided example of three": {
			[]int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			3,
			241861950,
		},
	}

	for _, tc := range cases {
		got, err := getMulAccounting(tc.account, tc.numOfparts)
		assert.Nil(t, err)
		assert.Equal(t, tc.want, got)
	}
}
