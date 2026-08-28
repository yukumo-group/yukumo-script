package characters_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/characters"
)

func TestNewCharacter(t *testing.T) {
	t.Parallel()
	img := "img.png"
	c := characters.NewCharacter("Yukumo", "f1", "desc", &img)
	if c.Name != "Yukumo" || c.PhontName != "f1" || c.Description != "desc" {
		t.Fatalf("unexpected character: %+v", c)
	}
	if !c.HasProfileImage() {
		t.Fatal("HasProfileImage want true")
	}
	c2 := characters.NewCharacter("A", "b", "c", nil)
	if c2.HasProfileImage() {
		t.Fatal("HasProfileImage want false")
	}
}

func TestCharactersCRUDAndJSON(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	store := characters.NewCharacters()
	store.SetTargetFile(dir, "characters.json")

	if err := store.ReadData(); err != nil {
		t.Fatalf("ReadData create: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "characters.json")); err != nil {
		t.Fatalf("characters.json missing: %v", err)
	}

	ch := characters.NewCharacter("Yukumo", "phont1", "hello", nil)
	if err := store.AddCharacter(ch); err != nil {
		t.Fatalf("AddCharacter: %v", err)
	}
	if err := store.AddCharacter(ch); err == nil {
		t.Fatal("AddCharacter duplicate: want error")
	}
	if err := store.AddCharacter(nil); err == nil {
		t.Fatal("AddCharacter nil: want error")
	}

	if err := store.SaveData(); err != nil {
		t.Fatalf("SaveData: %v", err)
	}

	loaded := characters.NewCharacters()
	loaded.SetTargetFile(dir, "characters.json")
	if err := loaded.ReadData(); err != nil {
		t.Fatalf("ReadData load: %v", err)
	}
	data := loaded.GetData()
	got, ok := data["Yukumo"]
	if !ok || got.Name != "Yukumo" || got.PhontName != "phont1" {
		t.Fatalf("loaded Yukumo = %+v, ok=%v", got, ok)
	}
	errDelete := loaded.DeleteCharacter("Yukumo")
	if errDelete != nil {
		t.Error(errDelete)
	}
	data = loaded.GetData()
	_, ok = data["Yukumo"]
	if ok {
		t.Errorf(
			"%s expected to get deleted",
			"Yukumo",
		)
	}
}

func TestCharacterDataClean(t *testing.T) {
	t.Parallel()
	newCharacter := characters.NewCharacters()
	newCharacter.Data = map[string]*characters.Character{
		"a": nil,
	}
	newCharacter.CleanData()
	characters := newCharacter.GetData()
	_, exists := characters["a"]
	if exists {
		t.Error("expect this one to get cleaned")
	}
}
