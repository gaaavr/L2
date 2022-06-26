package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	beforeTime := time.Now()
	TableTest := []struct {
		testName string
		t        time.Time
		expected bool
	}{
		{
			"first test : time before",
			beforeTime,
			false,
		},
		{
			"sec test : time after",
			time.Now().Add(1 * time.Minute),
			false,
		},
	}

	for _, testCase := range TableTest {
		time, err := getCurrentTime()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, time == testCase.t, testCase.expected)
	}
}
