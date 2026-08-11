package syncutils_test

import (
	"slices"
	"sync"
	"testing"

	"github.com/yukumo-group/yukumo-script/pkg/utils/syncutils"
)

func TestMapSetGetKeys(t *testing.T) {
	t.Parallel()
	m := syncutils.NewMap()

	if _, ok := m.GetValue("missing"); ok {
		t.Fatal("GetValue(missing) should not exist")
	}

	m.SetKV("a", "1")
	m.SetKV("b", "2")

	if v, ok := m.GetValue("a"); !ok || v != "1" {
		t.Fatalf("GetValue(a) = %q, %v; want 1, true", v, ok)
	}
	if v, ok := m.GetValue("b"); !ok || v != "2" {
		t.Fatalf("GetValue(b) = %q, %v; want 2, true", v, ok)
	}

	keys := m.GetAllKeys()
	if len(keys) != 2 {
		t.Fatalf("GetAllKeys len = %d, want 2", len(keys))
	}
	slices.Sort(keys)
	if !slices.Equal(keys, []string{"a", "b"}) {
		t.Fatalf("GetAllKeys = %v, want [a b]", keys)
	}
}

func TestMapConcurrent(t *testing.T) {
	t.Parallel()
	m := syncutils.NewMap()
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := string(rune('A' + i%26))
			m.SetKV(key, key)
			_, _ = m.GetValue(key)
			_ = m.GetAllKeys()
		}(i)
	}
	wg.Wait()
}
