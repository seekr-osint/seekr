package api
// useful functions in general
import (
  "sort"
)
func SortMapKeys[T comparable](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

