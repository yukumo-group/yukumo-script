package all2jap_test

import (
	"testing"
	"unicode"

	"github.com/dlclark/regexp2"

	"github.com/yukumo-group/yukumo-script/pkg/utils/language/all2jap"
)

var nonKatakana = regexp2.MustCompile(`[^\p{Katakana}\u30FC]+`, 0)

func assertOnlyKatakana(t *testing.T, label, s string) {
	t.Helper()
	if s == "" {
		t.Fatalf("%s: expected non-empty katakana output", label)
	}
	matchRes, err := nonKatakana.MatchString(s)
	if err != nil {
		t.Error(err)
	}
	if matchRes {
		t.Fatalf("%s: got non-katakana runes in %q", label, s)
	}
	for _, r := range s {
		if !unicode.In(r, unicode.Katakana) && (string(r) != "ー") {
			t.Fatalf("%s: rune %q is not katakana in %q", label, r, s)
		}
	}
}

func TestEngToKana(t *testing.T) {
	t.Parallel()
	got, err := all2jap.EngToKana("hello")
	if err != nil {
		t.Error(err)
	}
	assertOnlyKatakana(t, "EngToKana(hello)", got)

	withNum, err := all2jap.EngToKana("a1")
	if err != nil {
		t.Error(err)
	}
	assertOnlyKatakana(t, "EngToKana(a1)", withNum)
}

func TestJPToKana(t *testing.T) {
	t.Parallel()
	got, err := all2jap.JPToKana("こんにちは")
	if err != nil {
		t.Error(err)
	}
	assertOnlyKatakana(t, "JPToKana", got)

	withNum, err := all2jap.JPToKana("あ1")
	if err != nil {
		t.Error(err)
	}
	assertOnlyKatakana(t, "JPToKana(あ1)", withNum)
}

func TestAllToKana(t *testing.T) {
	t.Parallel()
	got, err := all2jap.AllToKana("Gopher")
	if err != nil {
		t.Error(err)
	}
	assertOnlyKatakana(t, "AllToKana(Gopher)", got)
}
