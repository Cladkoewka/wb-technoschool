package main

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input string
		expected string
		wantErr bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"a0", "", true},     
		{"3abc", "", true},   // starts with digit
		{"a10", "", true},		// digit more than 9
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\45\\", "", true}, // ends with \
		{"\\qwe\\45", "qwe44444", false},
	}

	for _, tt := range tests {
		got, err := Unpack(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("Unpack(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if got != tt.expected {
			t.Errorf("Unpack(%q) = %q, want %q", tt.input, got, tt.expected)
		}
		t.Logf("Test case input=%q output=%q err=%v", tt.input, got, err)
	}
}