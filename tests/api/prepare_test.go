package api_test

import (
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/internal/phontsmanager"
	"github.com/yukumo-group/yukumo-script/pkg/api"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

func TestPrepareGenerateByPhont(t *testing.T) {
	phont := "prep_phont_" + filepath.Base(t.TempDir())
	phontsmanager.PhontNameToFileName.SetKV(
		phont,
		phont+".phont",
	)

	taskName := "prep_task_" + filepath.Base(t.TempDir())
	task, err := api.PrepareGenerateByPhont(
		api.NewGenerateByPhontParams(
			taskName,
			"hello",
			language.English.ToInt(),
			100,
			phont,
		),
	)
	if err != nil {
		t.Fatalf("PrepareGenerateByPhont: %v", err)
	}
	if task.TaskName != taskName || task.PhontName == nil || *task.PhontName != phont {
		t.Fatalf("unexpected task: %+v", task)
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

	_, err := api.PrepareGenerateByPhont(
		api.NewGenerateByPhontParams(
			existing,
			"hi",
			language.English.ToInt(),
			100,
			phont,
		),
	)
	if err == nil {
		t.Fatal("duplicate task: want error")
	}

	_, err = api.PrepareGenerateByPhont(
		api.NewGenerateByPhontParams(
			"fresh_"+filepath.Base(dir),
			"hi",
			language.English.ToInt(),
			100,
			"no_such_phont_xyz",
		),
	)
	if err == nil {
		t.Fatal("missing phont: want error")
	}
}
