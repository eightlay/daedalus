package daedalus

import (
	"fmt"
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

func combine_errors(base_error error, occured_errors []error) error {
	if len(occured_errors) == 0 {
		return base_error
	}

	for _, step_err := range occured_errors {
		base_error = fmt.Errorf("%w\n%w", base_error, step_err)
	}

	return base_error
}

func prepend_to_error(prefix string, err error) error {
	return fmt.Errorf("%s %w", prefix, err)
}
