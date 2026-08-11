package api_test

import (
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/api"
	"github.com/yukumo-group/yukumo-script/pkg/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/language"
	"github.com/yukumo-group/yukumo-script/pkg/phontsmanager"
)

func TestPrepareGenerateByPhont(t *testing.T) {
	phont := "prep_phont_" + filepath.Base(t.TempDir())
	phontsmanager.PhontNameToFileName.SetKV(phont, phont+".phont")

	taskName := "prep_task_" + filepath.Base(t.TempDir())
	task, err := api.PrepareGenerateByPhont(api.GenerateByPhontParams{
		TaskName:  taskName,
		Text:      "hello",
		Language:  language.English.ToInt(),
		Speed:     100,
		PhontName: phont,
	})
	if err != nil {
		t.Fatalf("PrepareGenerateByPhont: %v", err)
	}
	if task.TaskName != taskName || task.PhontName == nil || *task.PhontName != phont {
		t.Fatalf("unexpected task: %+v", task)
	}
	if task.Text == "" || task.Text == "hello" {
		t.Fatalf("expected converted kana text, got %q", task.Text)
	}
	if task.Speed != 100 {
		t.Fatalf("Speed = %d, want 100", task.Speed)
	}
}

func TestPrepareGenerateByPhontErrors(t *testing.T) {
	phont := "prep_err_phont_" + filepath.Base(t.TempDir())
	phontsmanager.PhontNameToFileName.SetKV(phont, phont+".phont")

	dir := t.TempDir()
	singlesentence.Manager.SetTargetFile(dir, "tasks.json")
	_ = singlesentence.Manager.ReadData()

	existing := "existing_" + filepath.Base(dir)
	if err := singlesentence.Manager.NewTask(existing, "meta.json"); err != nil {
		t.Fatalf("seed NewTask: %v", err)
	}

	_, err := api.PrepareGenerateByPhont(api.GenerateByPhontParams{
		TaskName:  existing,
		Text:      "hi",
		Language:  language.English.ToInt(),
		Speed:     100,
		PhontName: phont,
	})
	if err == nil {
		t.Fatal("duplicate task: want error")
	}

	_, err = api.PrepareGenerateByPhont(api.GenerateByPhontParams{
		TaskName:  "fresh_" + filepath.Base(dir),
		Text:      "hi",
		Language:  language.English.ToInt(),
		Speed:     100,
		PhontName: "no_such_phont_xyz",
	})
	if err == nil {
		t.Fatal("missing phont: want error")
	}

	_, err = api.PrepareGenerateByPhont(api.GenerateByPhontParams{
		TaskName:  "lang_" + filepath.Base(dir),
		Text:      "hi",
		Language:  language.Chinese.ToInt(),
		Speed:     100,
		PhontName: phont,
	})
	if err == nil {
		t.Fatal("unsupported language: want error")
	}
}
