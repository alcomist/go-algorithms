package string

import (
	"reflect"
	"testing"
)

func TestKmp(t *testing.T) {

	k := NewKmp("aabaabacbba")

	var tests = []struct {
		input string
		want  []int
	}{
		{"aa", []int{0, 3}},
	}

	for _, test := range tests {
		if got := k.Search(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("kmp.Search(%q) = %v (WANT:%v)", test.input, got, test.want)
		}
	}
}
