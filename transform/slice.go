// Copyright 2022 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform

import (
	"github.com/juju/errors"
)

// Slice transforms a slice of one type to an equal length slice of another,
// by applying the input transformation function to each member.
func Slice[F any, T any](from []F, transform func(F) T) []T {
	to := make([]T, len(from))
	for i, oneFrom := range from {
		to[i] = transform(oneFrom)
	}
	return to
}

// SliceOrErr transforms a slice from one type to an equal length slice of another
// by mapping the input transformation function to each member.
// This differs from Slice in that the transform function can returns an error.
// If an error is encountered, the mapping will be cancelled and the error returned
func SliceOrErr[F any, T any](from []F, transform func(F) (T, error)) ([]T, error) {
	to := make([]T, len(from))
	for i, oneFrom := range from {
		var err error
		to[i], err = transform(oneFrom)
		if err != nil {
			return nil, errors.Annotatef(err, "error encountered transforming slice at index %d", i)
		}
	}
	return to, nil
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
