// Copyright 2022 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform

// Slice transforms a slice of one type to an equal length slice of another,
// by applying the input transformation function to each member.
func Slice[F any, T any](from []F, transform func(F) T) []T {
	to := make([]T, len(from))
	for i, oneFrom := range from {
		to[i] = transform(oneFrom)
	}
	return to
}

// SliceToMap transforms a slice of one type to an equal length
// map with values from the slice, keyed by values indicated by
// the input transformation function.
func SliceToMap[F any, K comparable, V any](from []F, transform func(F) (K, V)) map[K]V {
	to := make(map[K]V, len(from))
	for _, oneFrom := range from {
		k, v := transform(oneFrom)
		to[k] = v
	}
	return to
}
