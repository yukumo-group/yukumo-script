package chorus_test

import (
	"testing"

	"github.com/yukumo-group/yukumo-script/internal/generator/tasks"
	"github.com/yukumo-group/yukumo-script/internal/generator/tasks/chorus"
)

func TestInterface(
	t *testing.T,
) {
	t.Parallel()
	var testTask interface{} = &chorus.Task{}
	_, ok := testTask.(tasks.Task)
	if !ok {
		t.Error("Empty task cannot suit to the task interface")
	}
}
