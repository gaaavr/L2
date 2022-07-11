package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	TableTest := []struct {
		testName string
		arr      []string
		expected map[string][]string
	}{
		{
			"first test: correct",
			[]string{"пятка", "слиток", "пятак", "тяпка", "столик"},
			map[string][]string{
				"пятка":  {"пятак", "тяпка"},
				"слиток": {"столик"},
			},
		},
		{
			"second test: different register",
			[]string{"пЯтКа", "сЛиток", "пяТАк", "ТЯПКА", "столик"},
			map[string][]string{
				"пятка":  {"пятак", "тяпка"},
				"слиток": {"столик"},
			},
		},
		{
			"third test: empty value",
			[]string{},
			map[string][]string{},
		},
	}

	for _, testCase := range TableTest {
		l := findAnagrams(testCase.arr)
		assert.Equal(t, testCase.expected, l)
	}
}
