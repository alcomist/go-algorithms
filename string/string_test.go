package string

import (
	"reflect"
	"testing"
)

func TestKmp(t *testing.T) {

	k := NewKmp()
	haystack := "aabaabacbba"

	var tests = []struct {
		needle string
		want   []int
	}{
		{"aa", []int{0, 3}},
	}

	for _, test := range tests {
		if got := k.Search(test.needle, haystack); !reflect.DeepEqual(got, test.want) {
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
		if got := k.Search(test.needle, haystack); !reflect.DeepEqual(got, test.want) {
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
		if got := k.Search(test.needle, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("kmp.Search(%q, %q) = %v (WANT:%v)", test.needle, haystack, got, test.want)
		}
	}
}

func TestRabinKarp(t *testing.T) {

	r := NewRK()

	haystack := "aabaabacbba"

	var tests = []struct {
		input string
		want  []int
	}{
		{"aa", []int{0, 3}},
	}

	for _, test := range tests {
		if got := r.Search(test.input, haystack); !reflect.DeepEqual(got, test.want) {
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
		if got := r.Search(test.input, haystack); !reflect.DeepEqual(got, test.want) {
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
		if got := r.Search(test.input, haystack); !reflect.DeepEqual(got, test.want) {
			t.Errorf("rabin_karp.Search(%q, %q) = %v (WANT:%v)", test.input, haystack, got, test.want)
		}
	}
}
