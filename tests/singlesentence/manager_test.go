package singlesentence_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/singlesentence"
)

func TestTaskManagerJSON(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	mgr := singlesentence.NewTaskManager()
	mgr.SetTargetFile(dir, "tasks.json")

	if err := mgr.ReadData(); err != nil {
		t.Fatalf("ReadData create: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "tasks.json")); err != nil {
		t.Fatalf("tasks.json missing: %v", err)
	}

	if mgr.HasTask("t1") {
		t.Fatal("HasTask before NewTask")
	}
	if err := mgr.NewTask("t1", "file1.json"); err != nil {
		t.Fatalf("NewTask: %v", err)
	}
	if !mgr.HasTask("t1") {
		t.Fatal("HasTask after NewTask")
	}
	if err := mgr.NewTask("t1", "file1.json"); err == nil {
		t.Fatal("NewTask duplicate: want error")
	}

	all := mgr.GetAllTasks()
	if all["t1"] != "file1.json" {
		t.Fatalf("GetAllTasks = %v", all)
	}

	if err := mgr.DeleteTask("t1"); err != nil {
		t.Fatalf("DeleteTask: %v", err)
	}
	if mgr.HasTask("t1") {
		t.Fatal("HasTask after DeleteTask")
	}
	if err := mgr.DeleteTask("t1"); err == nil {
		t.Fatal("DeleteTask missing: want error")
	}

	// Reload from disk after save cycle.
	mgr2 := singlesentence.NewTaskManager()
	mgr2.SetTargetFile(dir, "tasks.json")
	if err := mgr2.ReadData(); err != nil {
		t.Fatalf("ReadData reload: %v", err)
	}
	if mgr2.HasTask("t1") {
		t.Fatal("reloaded manager still has deleted task")
	}
	characterID := "Remilia Scarlet"
	// Get file path
	task, err := singlesentence.NewSingleSentenceTask(
		"Hello, World",
		&characterID,
		nil,
		100,
		"a",
		1,
	)
	if err != nil {
		t.Error(err)
	}
	tmpDir := t.TempDir()
	filePath, err := task.SaveFile(tmpDir)
	if err != nil {
		t.Error(err)
	}
	err = mgr2.NewTask(
		task.TaskName,
		filePath,
	)
	if err != nil {
		t.Error(err)
	}
	newTask, err := mgr2.GetTask(
		task.TaskName,
	)
	if err != nil {
		t.Error(err)
	}
	if newTask.TaskName != task.TaskName {
		t.Errorf(
			"Expected task name %s, got %s",
			task.TaskName,
			newTask.TaskName,
		)
	}
}
