// Copyright 2023 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform

// MapToSlice is responsible for flattening a map of key values into a
// contiguous slice of key value pares from the map.
func MapToSlice[K comparable, V, T any](from map[K]V, transform func(K, V) []T) []T {
	to := make([]T, 0, len(from)*2)
	for k, v := range from {
		t := transform(k, v)
		to = append(to, t...)
	}
	return to
}
