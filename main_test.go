package main

import (
	"image/color"
	"reflect"
	"testing"
)

type ColorTest struct {
	input  string
	result color.RGBA
	error  bool
}

// Test the hex-string to colors.RGBA
// parser
func TestColor(t *testing.T) {
	vars := map[string]ColorTest{
		"base-read": {
			input: "abcdef",
			result: color.RGBA{
				R: 171,
				G: 205,
				B: 239,
				A: 255,
			},
			error: false,
		},
		"short-read": {
			input: "123",
			result: color.RGBA{
				R: 18,
				G: 49,
				B: 35,
				A: 255,
			},
			error: false,
		},
		"error-invalid": {
			input: "xyz",
			error: true,
		},
		"error-no-read": {
			input: "abcd",
			error: true,
		},
		"cutoff": {
			input: "ffffff",
			result: color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			},
			error: false,
		},
	}

	input := map[string]string{}

	for key, value := range vars {
		input[key] = value.input
		res, err := readColor(&input, key)
		if value.error && err == nil {
			t.Errorf("Could not properly detect %s as an error.", value.input)
		} else if !value.error && !reflect.DeepEqual(res, value.result) {
			t.Errorf("Could not properly detect color %s (%v) as %v.", value.input, res, value.result)
		}
	}
}
