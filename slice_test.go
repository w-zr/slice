// Copyright 2022 Wei Ziran
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdjacentFind(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		expected int
	}{
		{
			desc:     "s is nil",
			expected: -1,
		},
		{
			desc:     "with adjacent elements",
			slice:    []int{0, 3, 1, 1},
			expected: 2,
		},
		{
			desc:     "without adjacent elements",
			slice:    []int{0, 1, -1, 3, -3, 5, -5},
			expected: -1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, AdjacentFind(c.slice))
	}
}

func TestAllOf(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		pred     func(e int) bool
		expected bool
	}{
		{
			desc:     "all of the numbers are odd",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			pred:     func(e int) bool { return e%2 == 1 },
			expected: true,
		},
		{
			desc:     "some of the numbers are negative",
			slice:    []int{0, 1, -1, 3, -3, 5, -5},
			pred:     func(e int) bool { return e < 0 },
			expected: false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, AllOf(c.slice, c.pred))
	}
}

func TestAnyOf(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		pred     func(e int) bool
		expected bool
	}{
		{
			desc:     "none of the numbers are even",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			pred:     func(e int) bool { return e%2 == 0 },
			expected: false,
		},
		{
			desc:     "some of the numbers are negative",
			slice:    []int{0, 1, -1, 3, -3, 5, -5},
			pred:     func(e int) bool { return e < 0 },
			expected: true,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, AnyOf(c.slice, c.pred))
	}
}

func TestNoneOf(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		pred     func(e int) bool
		expected bool
	}{
		{
			desc:     "none of the numbers are even",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			pred:     func(e int) bool { return e%2 == 0 },
			expected: true,
		},
		{
			desc:     "some of the numbers are negative",
			slice:    []int{0, 1, -1, 3, -3, 5, -5},
			pred:     func(e int) bool { return e < 0 },
			expected: false,
		},
		{
			desc:     "none of the numbers are negative",
			slice:    []int{1, 2, 4, 8, 16, 32, 64, 128},
			pred:     func(e int) bool { return e < 0 },
			expected: true,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, NoneOf(c.slice, c.pred))
	}
}

func TestFindIf(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		pred     func(e int) bool
		expected int
	}{
		{
			desc:     "none of the numbers are even",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			pred:     func(e int) bool { return e%2 == 0 },
			expected: -1,
		},
		{
			desc:     "some of the numbers are negative",
			slice:    []int{0, 1, -1, 3, -3, 5, -5},
			pred:     func(e int) bool { return e < 0 },
			expected: 2,
		},
		{
			desc:     "none of the numbers are negative",
			slice:    []int{1, 2, 4, 8, 16, 32, 64, 128},
			pred:     func(e int) bool { return e < 0 },
			expected: -1,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, FindIf(c.slice, c.pred))
	}
}

func TestForEach(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		fn       func(e *int)
		expected []int
	}{
		{
			desc:     "minus one",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			fn:       func(e *int) { *e -= 1 },
			expected: []int{2, 4, 6, 10, 12, 16, 18, 22},
		},
		{
			desc:     "negative",
			slice:    []int{0, 1, 3, 5},
			fn:       func(e *int) { *e = -*e },
			expected: []int{0, -1, -3, -5},
		},
		{
			desc:     "over two",
			slice:    []int{1, 2, 4, 8, 16, 32, 64, 128},
			fn:       func(e *int) { *e /= 2 },
			expected: []int{0, 1, 2, 4, 8, 16, 32, 64},
		},
	}

	for _, c := range cases {
		ForEach(c.slice, c.fn)
		assert.Equal(t, c.expected, c.slice)
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		v        int
		expected int
	}{
		{
			desc:     "7",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			v:        7,
			expected: 2,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, Find(c.slice, c.v))
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		v        int
		expected int
	}{
		{
			desc:     "three ones",
			slice:    []int{1, 1, 2, 3, 4, 2, 1},
			v:        1,
			expected: 3,
		},
	}
	for _, c := range cases {
		assert.Equal(t, c.expected, Count(c.slice, c.v))
	}
}

func TestCountIf(t *testing.T) {
	cases := []struct {
		desc     string
		slice    []int
		f        func(int) bool
		expected int
	}{
		{
			desc:     "8 odd numbers",
			slice:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			f:        func(i int) bool { return i%2 == 1 },
			expected: 8,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, CountIf(c.slice, c.f))
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		desc string
		s1   []int
		s2   []int
		e    bool
	}{
		{
			desc: "equal",
			s1:   []int{3, 5, 7, 11, 13, 17, 19, 23},
			s2:   []int{3, 5, 7, 11, 13, 17, 19, 23},
			e:    true,
		},
		{
			desc: "unequal",
			s1:   []int{3, 5, 7, 11, 13, 17, 19, 23},
			s2:   []int{3, 5, 7, 11, 13, 17, 19},
			e:    false,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.e, Equal(c.s1, c.s2))
	}
}

func TestEqualIf(t *testing.T) {
	cases := []struct {
		desc string
		s1   []int
		s2   []int
		f    func(int) bool
		e    bool
	}{
		{
			desc: "all are not even",
			s1:   []int{3, 5, 7, 11, 13, 17, 19, 23},
			s2:   []int{3, 9, 7, 11, 13, 17, 99, 23},
			f: func(i int) bool {
				return i%2 == 0
			},
			e: true,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.e, EqualIf(c.s1, c.s2, c.f))
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		desc string
		s    []int
		v    int
		e    bool
	}{
		{
			desc: "all are not even",
			s:    []int{3, 5, 7, 11, 13, 17, 19, 23},
			v:    11,
			e:    true,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.e, Contains(c.s, c.v))
	}
}

func TestMax(t *testing.T) {
	assert.Equal(t, 5, Max(3, 5, func(x, y int) bool { return x > y }))
	assert.Equal(t, "abc", Max("abc", "a", func(x, y string) bool { return x > y }))
}
