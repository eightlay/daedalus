package daedalus

import (
	"sort"
)

func sort_map_keys(m interface{}) (keyList []int) {
	switch m := m.(type) {
	case map[int]*stage:
		for k := range m {
			keyList = append(keyList, k)
		}
	case map[int]Step:
		for k := range m {
			keyList = append(keyList, k)
		}
	default:
		panic("unknown map type")
	}

	sort.Ints(keyList)
	return
}
