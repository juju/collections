// Copyright 2023 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package transform_test

import (
	"fmt"

	"golang.org/x/exp/slices"
	. "gopkg.in/check.v1"

	"github.com/juju/collections/transform"
)

type mapSuite struct{}

var _ = Suite(mapSuite{})

func ExampleMapToSlice() {
	peopleStatus := map[string]string{
		"wallyworld": "peachy",
		"bob":        "happy",
	}

	flat := transform.MapToSlice(peopleStatus, func(k, v string) []string {
		return []string{k, v}
	})
	fmt.Println(flat)

	// Output:
	// [wallyworld peachy bob happy]
}

func (mapSuite) TestEmptyMapToSlice(c *C) {
	m := map[string]string{}
	to := transform.MapToSlice(m, func(k, v string) []any { return []any{k, v} })
	c.Assert(len(to), Equals, 0)
}

func (mapSuite) TestMapToSlice(c *C) {
	m := map[string]string{
		"a": "b",
		"c": "d",
	}
	to := transform.MapToSlice(m, func(k, v string) []string { return []string{k, v} })
	slices.Sort(to)
	c.Assert(to, DeepEquals, []string{"a", "b", "c", "d"})
}

func (mapSuite) TestEmptyMapTransformEmpty(c *C) {
	m := map[string]string{}

	to := transform.Map(m, func(k, v string) (string, any) {
		return k, v
	})

	c.Assert(len(to), Equals, 0)
}

func (mapSuite) TestEmptyMapTransform(c *C) {
	m := map[string]string{
		"one":   "two",
		"three": "four",
	}

	to := transform.Map(m, func(k, v string) (string, int) {
		if v == "two" {
			return k, 2
		}
		return k, 4
	})

	c.Assert(to, DeepEquals, map[string]int{
		"one":   2,
		"three": 4,
	})
}
