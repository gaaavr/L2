package main

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	TableTest := []struct {
		testName       string
		originalString string
		expected       string
		err            error
	}{
		{
			"first test: default string",
			"a4bc2d5e",
			"aaaabccddddde",
			nil,
		},
		{
			"second test: default string",
			"abcd",
			"abcd",
			nil,
		},
		{
			"third test: first number",
			"45",
			"",
			fmt.Errorf("(некорректная строка)"),
		},
		{
			"fourth test: empty string",
			"",
			"",
			nil,
		},
		{
			"fifth test: escape sequence",
			"qwe\\4\\5",
			"qwe45",
			nil,
		},
		{
			"sixth test: escape sequence",
			"qwe\\45",
			"qwe44444",
			nil,
		},
		{
			"seventh test: escape sequence",
			"qwe\\\\5",
			"qwe\\\\\\\\\\",
			nil,
		},
	}

	for _, testCase := range TableTest {
		result, err := unpackingString(testCase.originalString)
		assert.Equal(t, testCase.expected, result)
		assert.Equal(t, testCase.err, err)
	}

}
