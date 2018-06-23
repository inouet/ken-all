package util

import (
	"testing"
)

func TestUniqIsUnique(t *testing.T) {

	cases := []struct {
		data     string
		expected bool
	}{
		{data: "1", expected: true},
		{data: "1", expected: false},
		{data: "2", expected: true},
		{data: "1", expected: false},
		{data: "3", expected: true},
	}

	uniq := NewUniq()

	for _, c := range cases {
		actual := uniq.IsUnique(c.data)
		if actual != c.expected {
			t.Errorf("want '%v', got '%v'\n", c.expected, actual)
		}
	}
}
