package string

import (
	"github.com/alcomist/go-algorithms/string/distance"
	"github.com/alcomist/go-algorithms/string/interleaving"
	"github.com/alcomist/go-algorithms/string/kmp"
	"github.com/alcomist/go-algorithms/string/rabin_karp"
	"reflect"
	"testing"
)

func TestKmp(t *testing.T) {

	haystack := "aabaabacbba"

	var tests = []struct {
		needle string
		want   []int
	}{
		{"aa", []int{0, 3}},
	}

	for _, test := range tests {
		if got := kmp.Search(test.needle, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("kmp.Search(%q, %q) = %v (WANT:%v)", test.needle, haystack, got, test.want)
		}
	}

	haystack = "토마토마토마한글한글한"

	tests = []struct {
		needle string
		want   []int
	}{
		{"마토", []int{1, 3}},
		{"글한", []int{7, 9}},
		{"영어", []int{}},
	}

	for _, test := range tests {
		if got := kmp.Search(test.needle, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("kmp.Search(%q, %q) = %v (WANT:%v)", test.needle, haystack, got, test.want)
		}
	}

	haystack = "你明天怎么去么去么去明天明"

	tests = []struct {
		needle string
		want   []int
	}{
		{"去么去", []int{5, 7}},
		{"去么", []int{5, 7}},
		{"明天", []int{1, 10}},
	}

	for _, test := range tests {
		if got := kmp.Search(test.needle, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("kmp.Search(%q, %q) = %v (WANT:%v)", test.needle, haystack, got, test.want)
		}
	}
}

func TestRabinKarp(t *testing.T) {

	haystack := "aabaabacbba"

	var tests = []struct {
		input string
		want  []int
	}{
		{"aa", []int{0, 3}},
	}

	for _, test := range tests {
		if got := rabin_karp.Search(test.input, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("rabin_karp.Search(%q, %q) = %v (WANT:%v)", test.input, haystack, got, test.want)
		}
	}

	haystack = "토마토마토마한글한글한"

	tests = []struct {
		input string
		want  []int
	}{
		{"마토", []int{1, 3}},
		{"글한", []int{7, 9}},
		{"영어", []int{}},
	}

	for _, test := range tests {
		if got := rabin_karp.Search(test.input, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("rabin_karp.Search(%q, %q) = %v (WANT:%v)", test.input, haystack, got, test.want)
		}
	}

	haystack = "你明天怎么去么去么去明天明"

	tests = []struct {
		input string
		want  []int
	}{
		{"去么去", []int{5, 7}},
		{"去么", []int{5, 7}},
		{"明天", []int{1, 10}},
	}

	for _, test := range tests {
		if got := rabin_karp.Search(test.input, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("rabin_karp.Search(%q, %q) = %v (WANT:%v)", test.input, haystack, got, test.want)
		}
	}
}

func TestDistance(t *testing.T) {

	var tests = []struct {
		s1, s2 string
		want   int
	}{
		{"sunday", "saturday", 3},
		{"", "", 0},
		{"cat", "car", 1},
		{"abc", "ab", 1},
	}

	for _, test := range tests {
		if got := distance.Distance(test.s1, test.s2); got != test.want {
			t.Errorf("distance.Distance(%q, %q) = %v (WANT:%v)", test.s1, test.s2, got, test.want)
		}
	}
}

func TestInterleaving(t *testing.T) {

	var tests = []struct {
		a, b, c string
		want    bool
	}{
		{"xyz", "abcd", "xabyczd", true},
		{"bcc", "bbca", "bbcbcac", true},
		{"cat", "tree", "tcraete", true},
		{"cat", "tree", "catrtee", true},
		{"cat", "tree", "cttaree", false},
	}

	for _, test := range tests {
		if got := interleaving.Interleaving(test.a, test.b, test.c); got != test.want {
			t.Errorf("interleaving.Interleaving(%q, %q, %q) = %v (WANT:%v)", test.a, test.b, test.c, got, test.want)
		}
	}
}
