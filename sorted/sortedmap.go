package sorted

type SortedMap struct {
	Keys   []string
	Values []int
}

func (sm *SortedMap) Set(key string, value int) {
	sm.Keys = append(sm.Keys, key)
	sm.Values = append(sm.Values, value)
}

func (sm *SortedMap) Get(key string) (int, bool) {
	for i, k := range sm.Keys {
		if k == key {
			return sm.Values[i], true
		}
	}
	return 0, false
}
