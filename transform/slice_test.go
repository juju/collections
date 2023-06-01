// Copyright 2022 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform_test

import (
	"github.com/juju/collections/transform"
	"github.com/juju/errors"
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

func (sliceSuite) TestSliceOrErrTransformationSucceeds(c *gc.C) {
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

	thisToThat := func(from this) (that, error) { return that{numero: from.number}, nil }

	to := []that{
		{numero: 1},
		{numero: 2},
	}

	res, err := transform.SliceOrErr(from, thisToThat)
	c.Assert(err, gc.IsNil)
	c.Assert(res, gc.DeepEquals, to)
}

func (sliceSuite) TestSliceOrErrTransformationErrors(c *gc.C) {
	type this struct {
		number int
	}

	type that struct {
		numero int
	}

	from := []this{
		{number: 1},
		{number: 0},
		{number: 2},
	}

	thisToThat := func(from this) (that, error) {
		if from.number == 0 {
			return that{}, errors.New("cannot transform 0")
		}
		return that{numero: from.number}, nil
	}

	_, err := transform.SliceOrErr(from, thisToThat)
	c.Assert(err, gc.ErrorMatches, "error encountered transforming slice at index 1: cannot transform 0")
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
