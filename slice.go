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

// AdjacentFind searches s for the first occurrence of two consecutive elements that match,
// and returns the index to the first of these two elements, or -1 if no such pair is found.
func AdjacentFind[T comparable](s []T) int {
	for i := range s {
		if i == len(s)-1 {
			break
		}
		if s[i] == s[i+1] {
			return i
		}
	}
	return -1
}

// AllOf returns true if f returns true for all the elements in the s
// or if s is empty, and false otherwise.
func AllOf[T any](s []T, f func(T) bool) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

// AnyOf returns true if f returns true for any of the elements in s,
// and false otherwise.
func AnyOf[T any](s []T, f func(T) bool) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

// Count returns the number of elements in s
// that compare equal to v.
func Count[T comparable](s []T, v T) (ret int) {
	for i := range s {
		if s[i] == v {
			ret++
		}
	}
	return ret
}

// CountIf returns the number of elements in s for which f is true.
func CountIf[T any](s []T, f func(T) bool) (ret int) {
	for i := range s {
		if f(s[i]) {
			ret++
		}
	}
	return ret
}

// Contains reports whether v is present in s.
func Contains[T comparable](s []T, v T) bool {
	for i := range s {
		if s[i] == v {
			return true
		}
	}
	return false
}

// Equal reports whether two slices are equal: the same length and all elements equal.
// If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order,
// and the comparison stops at the first unequal pair.
func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// EqualIf reports whether two slices are equal:
// the same length and f(element) equal.
// If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order,
// and the comparison stops at the first unequal pair.
func EqualIf[T comparable](s1, s2 []T, f func(T) bool) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if f(s1[i]) != f(s2[i]) {
			return false
		}
	}

	return true
}

// Find returns the index to the first element in s that compares equal to v.
// If no such element is found, the function returns -1.
func Find[T comparable](s []T, v T) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}
	return -1
}

// FindIf returns the index to the first element in s for
// which f returns true. If no such element is found, the function returns -1.
func FindIf[T any](s []T, f func(T) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// ForEach applies function f to each of the element in s.
func ForEach[T any](s []T, f func(*T)) {
	for i := range s {
		f(&s[i])
	}
}

// Max returns y if f returns true. Otherwise, returns x.
func Max[T any](x, y T, f func(T, T) bool) T {
	if f(x, y) {
		return x
	}
	return y
}

// NoneOf returns true if f returns false for all the elements in s
// or if s is empty, and false otherwise.
func NoneOf[T any](s []T, f func(T) bool) bool {
	return !AnyOf(s, f)
}
