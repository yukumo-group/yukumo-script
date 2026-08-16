package audio_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/audio"
	"github.com/yukumo-group/yukumo-script/pkg/utils/audio/edit"
)

func TestUpdateChannelNumberTo2(t *testing.T) {
	t.Parallel()
	in := []int16{1, 2, 3}
	got := edit.UpdateChannelNumberTo2(in)
	want := []int16{1, 1, 2, 2, 3, 3}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got[%d] = %d, want %d (full %v)", i, got[i], want[i], got)
		}
	}
}

func TestResampleWAV(t *testing.T) {
	t.Parallel()
	// Mono tone samples at 8kHz → 16kHz.
	in := make([]int16, 80)
	for i := range in {
		in[i] = int16(i)
	}
	out, err := edit.ResampleWAV(in, 1, 8000, 16000)
	if err != nil {
		t.Fatalf("ResampleWAV: %v", err)
	}
	if len(out) == 0 {
		t.Fatal("ResampleWAV returned empty data")
	}
}

func TestConvertAllWAVCopy(t *testing.T) {
	t.Parallel()
	srcDir := t.TempDir()
	dstDir := t.TempDir()
	src := filepath.Join(srcDir, "sample.wav")
	if err := os.WriteFile(src, []byte("RIFF-fake"), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := audio.ConvertAll(src, dstDir, "out", audio.WAV); err != nil {
		t.Fatalf("ConvertAll WAV: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dstDir, "out.wav"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "RIFF-fake" {
		t.Fatalf("copied content = %q", got)
	}

	if err := audio.ConvertAll(src, dstDir, "out", audio.FLAC); err == nil {
		t.Fatal("ConvertAll FLAC: want unsupported error")
	}
}
