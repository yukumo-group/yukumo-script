package singlesentence_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/generator/tasks/singlesentence"
)

func TestNewSingleSentenceTaskRequiresVoice(t *testing.T) {
	t.Parallel()
	_, err := singlesentence.NewSingleSentenceTask("text", nil, nil, 100, "t")
	if err == nil {
		t.Fatal("want error when both characterID and phontName are nil")
	}
}

func TestTaskSaveFileRoundTrip(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	phont := "f1"
	task, err := singlesentence.NewSingleSentenceTask("カナ", nil, &phont, 100, "hello")
	if err != nil {
		t.Fatalf("NewSingleSentenceTask: %v", err)
	}
	if task.ID == "" || task.TaskName != "hello" || task.Text != "カナ" || task.Speed != 100 {
		t.Fatalf("unexpected task: %+v", task)
	}
	if task.ResultFile != nil {
		t.Fatal("ResultFile should be nil before generate")
	}

	path, err := task.SaveFile(dir)
	if err != nil {
		t.Fatalf("SaveFile: %v", err)
	}

	loaded, err := singlesentence.NewSingleSentenceTaskFromFile(path)
	if err != nil {
		t.Fatalf("NewSingleSentenceTaskFromFile: %v", err)
	}
	if loaded.ID != task.ID || loaded.TaskName != task.TaskName || loaded.Text != task.Text || loaded.Speed != task.Speed {
		t.Fatalf("loaded %+v, want %+v", loaded, task)
	}
	if loaded.PhontName == nil || *loaded.PhontName != phont {
		t.Fatalf("loaded PhontName = %v, want %s", loaded.PhontName, phont)
	}
}
