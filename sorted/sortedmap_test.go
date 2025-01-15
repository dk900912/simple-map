package sorted

import "testing"

func TestSortedMap(t *testing.T) {
	sm := &SortedMap{}

	// 测试 Set 方法
	sm.Set("key1", 1)
	sm.Set("key2", 2)

	// 测试 Get 方法
	value, ok := sm.Get("key1")
	if !ok || value != 1 {
		t.Errorf("Expected value 1 for key1, got %d, ok: %v", value, ok)
	}

	value, ok = sm.Get("key2")
	if !ok || value != 2 {
		t.Errorf("Expected value 2 for key2, got %d, ok: %v", value, ok)
	}

	// 测试不存在的键
	value, ok = sm.Get("key3")
	if ok {
		t.Errorf("Expected key3 to not exist, but got value: %d", value)
	}

	// 测试插入顺序
	sm.Set("key3", 3)
	expectedKeys := []string{"key1", "key2", "key3"}
	for i, key := range expectedKeys {
		if sm.Keys[i] != key {
			t.Errorf("Expected key %s at index %d, got %s", key, i, sm.Keys[i])
		}
	}
}
