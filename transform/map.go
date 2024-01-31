// Copyright 2023 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform

// Map transforms a map of one type to a map of another, by applying the input
// transformation function to each key value pair.
func Map[K comparable, V any, Kn comparable, Vn any](
	from map[K]V,
	transform func(K, V) (Kn, Vn),
) map[Kn]Vn {
	to := make(map[Kn]Vn, len(from))
	for k, v := range from {
		kn, vn := transform(k, v)
		to[kn] = vn
	}
	return to
}

// MapToSlice is responsible for flattening a map of key values into a
// contiguous slice of key value pairs from the map.
func MapToSlice[K comparable, V, T any](from map[K]V, transform func(K, V) []T) []T {
	to := make([]T, 0, len(from)*2)
	for k, v := range from {
		t := transform(k, v)
		to = append(to, t...)
	}
	return to
}
