package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMakeSheets(t *testing.T) {
	type TD struct {
		in1    [][]bool
		in2    []string
		expect [][]string
	}
	tds := []TD{
		TD{
			in1: [][]bool{
				[]bool{true, false},
				[]bool{false, true},
			},
			in2: []string{"1", "2"},
			expect: [][]string{
				[]string{"1", ""},
				[]string{"", "2"},
			},
		},
		TD{
			in1: [][]bool{
				[]bool{true, true, true},
			},
			in2: []string{"1", "a", "tanaka"},
			expect: [][]string{
				[]string{"1", "a", "tanaka"},
			},
		},
		TD{
			in1: [][]bool{
				[]bool{false, false},
			},
			in2: []string{"1", "a", "tanaka"},
			expect: [][]string{
				[]string{"", ""},
			},
		},
		TD{
			in1: [][]bool{
				[]bool{true, true},
			},
			in2: []string{"1"},
			expect: [][]string{
				[]string{"1", ""},
			},
		},
	}

	for _, v := range tds {
		in1, in2, expect := v.in1, v.in2, v.expect
		assert.Equal(t, expect, MakeSheets(in1, in2), "正常系")
	}
}
