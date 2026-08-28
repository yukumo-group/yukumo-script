package singlesentence_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
	"github.com/yukumo-group/yukumo-script/pkg/utils/language"
)

func TestNewSingleSentenceTaskRequiresVoice(t *testing.T) {
	t.Parallel()
	_, err := singlesentence.NewSingleSentenceTask("text", nil, nil, 100, "t", language.Japanese, nil)
	if err == nil {
		t.Fatal("want error when both characterID and phontName are nil")
	}
}

func TestTaskSaveFileRoundTrip(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	phont := "f1"
	task, err := singlesentence.NewSingleSentenceTask("カナ", nil, &phont, 100, "hello", language.Japanese, nil)
	if err != nil {
		t.Fatalf("NewSingleSentenceTask: %v", err)
	}
	if task.ID == "" || task.GetTaskName() != "hello" || task.Text != "カナ" || task.Speed != 100 {
		t.Fatalf("unexpected task: %+v", task)
	}
	if task.GetResultFile() != nil {
		t.Fatal("ResultFile should be nil before generate")
	}

	path, err := task.SaveFile(dir)
	if err != nil {
		t.Fatalf("SaveFile: %v", err)
	}

	loaded, err := singlesentence.NewSingleSentenceTaskFromFile(path, nil)
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

func TestInterface(t *testing.T) {
	var testTask interface{} = &singlesentence.Task{}
	_, ok := testTask.(tasks.Task)
	if !ok {
		t.Error("singlesentence.Task does not support task interface")
	}
}

func TestGenerateFileName(t *testing.T) {
	t.Parallel()
	cid := "Reimu"
	testTask, err := singlesentence.NewSingleSentenceTask(
		"text",
		&cid,
		nil,
		100,
		"t",
		language.Japanese,
		nil,
	)
	if err != nil {
		t.Error(err)
	}
	chan1 := make(chan string)
	chan2 := make(chan string)
	var result1 string
	var result2 string
	tmpDir := t.TempDir()
	func1 := func() {
		chan1 <- testTask.GenerateWavFileName(
			tmpDir,
		)
	}
	func2 := func() {
		chan2 <- testTask.GenerateWavFileName(
			tmpDir,
		)
	}
	go func1()
	go func2()
	for i := 0; i < 2; i++ {
		select {
		case result1 = <-chan1:
		case result2 = <-chan2:
		}
	}
	if result1 == result2 {
		t.Errorf(
			"expected two result %s and %s not to be the same",
			result1,
			result2,
		)
	}
}
