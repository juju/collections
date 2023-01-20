// Copyright 2022 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform_test

import (
	"github.com/juju/collections/transform"
	"github.com/juju/testing"
	gc "gopkg.in/check.v1"
)

type sliceSuite struct {
	testing.IsolationSuite
}

var _ = gc.Suite(sliceSuite{})

func (sliceSuite) TestSliceTransformation(c *gc.C) {
	type this struct {
		number int
	}

	type that struct {
		numero int
	}

	from := []this{
		{number: 1},
		{number: 2},
	}

	thisToThat := func(from this) that { return that{numero: from.number} }

	to := []that{
		{numero: 1},
		{numero: 2},
	}

	c.Assert(transform.Slice(from, thisToThat), gc.DeepEquals, to)
}

func (sliceSuite) TestSliceToMapTransformation(c *gc.C) {
	type this struct {
		number int
	}

	type that struct {
		numero int
	}

	from := []this{
		{number: 1},
		{number: 2},
	}

	thisToThat := func(from this) (int, that) { return from.number, that{numero: from.number} }

	to := map[int]that{
		1: {numero: 1},
		2: {numero: 2},
	}

	c.Assert(transform.SliceToMap(from, thisToThat), gc.DeepEquals, to)
}
