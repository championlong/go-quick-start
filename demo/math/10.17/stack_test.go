package stack

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_isValid(t *testing.T) {
	testCases := []struct {
		str  string
		want bool
	}{
		{
			str:  "}{",
			want: false,
		},
		{
			str:  "{}",
			want: true,
		},
	}

	for i, tt := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := isValid(tt.str)
			assert.EqualValues(t, tt.want, result)
		})
	}
}
