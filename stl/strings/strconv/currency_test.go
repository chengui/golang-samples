package strconv

import "testing"

func TestFormatCurrency(t *testing.T) {
	var tests = []struct {
		num  Currency
		want string
	}{
		{1, "1"},
		{1234, "1,234"},
		{123456, "123,456"},
		{12345678, "12,345,678"},
		{123456789, "123,456,789"},
		{-1234, "-1,234"},
		{-123456, "-123,456"},
	}
	for _, test := range tests {
		if got := FormatCurrency(test.num); got != test.want {
			t.Errorf("FormatCurrency(%d) got %s != %s", test.num, got, test.want)
		} else {
			t.Log(test.num, got)
		}
	}
}
