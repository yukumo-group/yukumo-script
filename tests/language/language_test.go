package language_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/language"
)

func TestToLanguage(t *testing.T) {
	t.Parallel()
	cases := []struct {
		idx  int
		want language.Language
	}{
		{0, language.Japanese},
		{1, language.English},
		{2, language.Chinese},
		{99, language.Japanese},
	}
	for _, tc := range cases {
		if got := language.ToLanguage(tc.idx); got != tc.want {
			t.Errorf("ToLanguage(%d) = %v, want %v", tc.idx, got, tc.want)
		}
	}
}

func TestLanguageToInt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		lang language.Language
		want int
	}{
		{language.Japanese, 0},
		{language.English, 1},
		{language.Chinese, 2},
		{language.Language(99), 0},
	}
	for _, tc := range cases {
		if got := tc.lang.ToInt(); got != tc.want {
			t.Errorf("%v.ToInt() = %d, want %d", tc.lang, got, tc.want)
		}
	}
}

func TestConvertText(t *testing.T) {
	t.Parallel()

	en, err := language.ConvertText("hello", language.English)
	if err != nil {
		t.Fatalf("ConvertText English: %v", err)
	}
	if en == "" {
		t.Fatal("ConvertText English returned empty string")
	}

	jp, err := language.ConvertText("こんにちは", language.Japanese)
	if err != nil {
		t.Fatalf("ConvertText Japanese: %v", err)
	}
	if jp == "" {
		t.Fatal("ConvertText Japanese returned empty string")
	}

	_, err = language.ConvertText("你好", language.Chinese)
	if err == nil {
		t.Fatal("ConvertText Chinese: want error, got nil")
	}
}
