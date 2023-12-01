package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDigit(t *testing.T) {
	tables := []struct {
		input    string
		expected int64
	}{
		{
			"1abc2", 12,
		},
		{
			"pqr3stu8vwx", 38,
		},
		{
			"a1b2c3d4e5f", 15,
		},
		{
			"treb7uchet", 77,
		},
		{
			"2njsevenszzsfltconesixhsflpbpd", 22,
		},
	}

	for _, table := range tables {
		x := findDigit(table.input)
		assert.Equal(t, table.expected, x, "they should be equal")
	}

}
