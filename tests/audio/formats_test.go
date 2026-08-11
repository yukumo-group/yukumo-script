package audio_test

import (
	"slices"
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
)

func TestToFormatAndToString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in   string
		want audio.Format
		str  string
	}{
		{"wav", audio.WAV, "wav"},
		{"mp3", audio.MP3, "mp3"},
		{"m4a", audio.AAC, "m4a"},
		{"flac", audio.FLAC, "flac"},
		{"unknown", audio.WAV, "wav"},
	}
	for _, tc := range cases {
		got := audio.ToFormat(tc.in)
		if got != tc.want {
			t.Errorf("ToFormat(%q) = %v, want %v", tc.in, got, tc.want)
		}
		if s := got.ToString(); s != tc.str {
			t.Errorf("%v.ToString() = %q, want %q", got, s, tc.str)
		}
	}
}

func TestGetAllFormats(t *testing.T) {
	t.Parallel()
	got := audio.GetAllFormats()
	want := []audio.Format{audio.WAV, audio.MP3, audio.AAC, audio.FLAC}
	if !slices.Equal(got, want) {
		t.Fatalf("GetAllFormats() = %v, want %v", got, want)
	}
}
