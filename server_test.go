//go:demo_tests

package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAddDays(t *testing.T) {
	var t_test Timestamp
	d := 1
	got := t_test.addDays(d)
	expected := Timestamp(int(1) + d*24*3600)
	assert.Equal(t, got, expected)

}

func TestAddHours(t *testing.T) {
	var t_test Timestamp
	h := 1
	got := t_test.addHours(h)
	expected := Timestamp(int(1) + h*3600)
	assert.Equal(t, got, expected)
}
