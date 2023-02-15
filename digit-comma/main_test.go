package main

import (
    "testing"
    "reflect"
)

func TestComma1(t *testing.T) {
	gots := []string{
		"12", "123", "1234", "12345", "123456", "1234567", "12345678",
		"-12", "-123", "-1234", "+12345", "-123456", "-1234567", "-12345678",
	}
	wants := []string{
		"12", "123", "1,234", "12,345", "123,456", "1,234,567", "12,345,678",
		"-12", "-123", "-1,234", "+12,345", "-123,456", "-1,234,567", "-12,345,678",
	}
	ln := len(gots)
	for i := 0; i < ln; i++ {
		got := Comma1(gots[i])
		want := wants[i]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get: \"%v\" ;Want: \"%v\"", got, want)
		}
	}
}

func TestComma2(t *testing.T) {
	gots := []string{
		"12", "123", "1234", "12345", "123456", "1234567", "12345678",
		"-12", "-123", "-1234", "+12345", "-123456", "-1234567", "-12345678",
	}
	wants := []string{
		"12", "123", "1,234", "12,345", "123,456", "1,234,567", "12,345,678",
		"-12", "-123", "-1,234", "+12,345", "-123,456", "-1,234,567", "-12,345,678",
	}
	ln := len(gots)
	for i := 0; i < ln; i++ {
		got := Comma1(gots[i])
		want := wants[i]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get: \"%v\" ;Want: \"%v\"", got, want)
		}
	}
}

func TestComma3(t *testing.T) {
	gots := []string{
		"12", "123", "1234", "12345", "123456", "1234567", "12345678",
		"-12", "-123", "-1234", "+12345", "-123456", "-1234567", "-12345678",
	}
	wants := []string{
		"12", "123", "1,234", "12,345", "123,456", "1,234,567", "12,345,678",
		"-12", "-123", "-1,234", "+12,345", "-123,456", "-1,234,567", "-12,345,678",
	}
	ln := len(gots)
	for i := 0; i < ln; i++ {
		got := Comma1(gots[i])
		want := wants[i]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get: \"%v\" ;Want: \"%v\"", got, want)
		}
	}
}

func TestComma4(t *testing.T) {
	gots := []string{
		"12", "123", "1234", "12345", "123456", "1234567", "12345678",
		"-12.1234", "-123.1234", "-1234.1234", "+12345.1234", "-123456.1234", "-1234567.1234", "-12345678.1234",
	}
	wants := []string{
		"12", "123", "1,234", "12,345", "123,456", "1,234,567", "12,345,678",
		"-12.1234", "-123.1234", "-1,234.1234", "+12,345.1234", "-123,456.1234", "-1,234,567.1234", "-12,345,678.1234",
	}
	ln := len(gots)
	for i := 0; i < ln; i++ {
		got := Comma4(gots[i])
		want := wants[i]
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Get: \"%v\" ;Want: \"%v\"", got, want)
		}
	}
}

func BenchmarkComma1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Comma1("-123456")
        Comma1("-12345678")
    }
}

func BenchmarkComma2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Comma2("-123456")
        Comma2("-12345678")
    }
}

func BenchmarkComma3(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Comma3("-123456")
        Comma3("-12345678")
    }
}
