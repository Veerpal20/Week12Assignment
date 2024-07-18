package main

import "testing"

// TestIsLeapYear tests the IsLeapYear function.
func TestIsLeapYear(t *testing.T) {
	cases := []struct {
		year int
		want bool
	}{
		{2000, true},
		{2004, true},
		{2018, false},
		{1900, false},
		{1600, true},
	}

	for _, c := range cases {
		got := IsLeapYear(c.year)
		if got != c.want {
			t.Errorf("IsLeapYear(%d) = %v, want %v", c.year, got, c.want)
		}
	}
}
