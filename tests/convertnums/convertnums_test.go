package convertnums_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/language/convertnums"
)

func TestConvertNumToEnglish(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"abc", "abc"},
		{"1", " one "},
		{"10", " one  zero "},
		{"a2b", "a two b"},
	}
	for _, tc := range cases {
		if got := convertnums.ConvertNumToEnglish(tc.in); got != tc.want {
			t.Errorf("ConvertNumToEnglish(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestConverNumToJP(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"abc", "abc"},
		{"1", " イチ "},
		{"10", " イチ  ゼロ "},
		{"a5b", "a ゴ b"},
	}
	for _, tc := range cases {
		if got := convertnums.ConverNumToJP(tc.in); got != tc.want {
			t.Errorf("ConverNumToJP(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}
