package phontsmanager_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
)

func TestInitializeAndGetPhontFile(t *testing.T) {
	dir := t.TempDir()
	phontName := "voice_test_" + filepath.Base(dir)
	if err := os.WriteFile(filepath.Join(dir, phontName+".phont"), []byte("phont"), 0644); err != nil {
		t.Fatalf("WriteFile phont: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("ignore"), 0644); err != nil {
		t.Fatalf("WriteFile txt: %v", err)
	}

	if err := phontsmanager.InitializePhontNameToFileName(dir); err != nil {
		t.Fatalf("InitializePhontNameToFileName: %v", err)
	}

	file, ok := phontsmanager.PhontNameToFileName.GetValue(phontName)
	if !ok || file != phontName+".phont" {
		t.Fatalf("map[%s] = %q, %v; want %s.phont", phontName, file, ok, phontName)
	}

	path, err := phontsmanager.GetPhontFile(dir, phontName)
	if err != nil {
		t.Fatalf("GetPhontFile: %v", err)
	}
	want := filepath.ToSlash(filepath.Join(dir, phontName+".phont"))
	if filepath.ToSlash(path) != want {
		t.Fatalf("GetPhontFile path = %q, want %q", path, want)
	}

	if _, err := phontsmanager.GetPhontFile(dir, "does-not-exist"); err == nil {
		t.Fatal("GetPhontFile missing: want error")
	}
}

func TestGetAllPhonts(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.phont"), []byte("x"), 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	entries, err := phontsmanager.GetAllPhonts(dir)
	if err != nil {
		t.Fatalf("GetAllPhonts: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("GetAllPhonts len = %d, want 1", len(entries))
	}
}
