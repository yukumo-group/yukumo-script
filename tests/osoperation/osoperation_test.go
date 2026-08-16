package osoperation_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/osoperation"
)

func TestParseWindowsPath(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, want string
	}{
		{`C:\foo\bar`, "C:/foo/bar"},
		{"already/unix", "already/unix"},
		{"", ""},
	}
	for _, tc := range cases {
		if got := osoperation.ParseWindowsPath(tc.in); got != tc.want {
			t.Errorf("ParseWindowsPath(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestGetNewFilePathAndSaveDataTo(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	path, err := osoperation.GetNewFilePath(dir, "note", "txt")
	if err != nil {
		t.Fatalf("GetNewFilePath: %v", err)
	}
	want := filepath.ToSlash(filepath.Join(dir, "note.txt"))
	if filepath.ToSlash(path) != want {
		t.Fatalf("path = %q, want %q", path, want)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("created file missing: %v", err)
	}

	// Existing path is returned without error.
	path2, err := osoperation.GetNewFilePath(dir, "note", "txt")
	if err != nil {
		t.Fatalf("GetNewFilePath existing: %v", err)
	}
	if path2 != path {
		t.Fatalf("existing path = %q, want %q", path2, path)
	}

	if err := osoperation.SaveDataTo(dir, "data", "bin", []byte("hello")); err != nil {
		t.Fatalf("SaveDataTo: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dir, "data.bin"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "hello" {
		t.Fatalf("content = %q, want hello", got)
	}
}

func TestCopyFile(t *testing.T) {
	t.Parallel()
	srcDir := t.TempDir()
	dstDir := t.TempDir()
	src := filepath.Join(srcDir, "src.txt")
	if err := os.WriteFile(src, []byte("payload"), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	if err := osoperation.CopyFile(src, dstDir, "copied", "txt"); err != nil {
		t.Fatalf("CopyFile: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dstDir, "copied.txt"))
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "payload" {
		t.Fatalf("copied = %q, want payload", got)
	}

	if err := osoperation.CopyFile(filepath.Join(srcDir, "missing.txt"), dstDir, "x", "txt"); err == nil {
		t.Fatal("CopyFile missing: want error")
	}
}

func TestImageSupported(
	t *testing.T,
) {
	suffix1, supported1 := osoperation.IsSuffixSupported(
		"abcdefg",
		osoperation.SupportedImageSuffix,
	)
	if supported1 {
		t.Errorf(
			"%s should not be supported. ",
			"abcdefg",
		)
	}
	if suffix1 != "" {
		t.Errorf(
			"expected %s, got %s",
			"",
			suffix1,
		)
	}
	suffix2, supported2 := osoperation.IsSuffixSupported(
		"abc.png",
		osoperation.SupportedImageSuffix,
	)
	if !supported2 {
		t.Errorf(
			"%s should be supported. ",
			"abc.png",
		)
	}
	if suffix2 != ".png" {
		t.Errorf(
			"expected %s, got %s",
			"png",
			suffix2,
		)
	}
	suffix3, supported3 := osoperation.IsSuffixSupported(
		"abc.go",
		osoperation.SupportedImageSuffix,
	)
	if supported3 {
		t.Errorf(
			"%s should not be supported. ",
			"abc.go",
		)
	}
	if suffix3 != ".go" {
		t.Errorf(
			"expected %s, got %s",
			".go",
			suffix3,
		)
	}
}
