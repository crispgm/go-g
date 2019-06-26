package string

// ExistsInSlice checks whether target exists in a string slice
func ExistsInSlice(target string, entries []string) bool {
	for _, item := range entries {
		if item == target {
			return true
		}
	}

	return false
}

// DedupSlice de-duplicates a string slice
func DedupSlice(entries []string) []string {
	var dedupKeys = make(map[string]bool, len(entries))
	var result []string
	for _, item := range entries {
		if _, ok := dedupKeys[item]; !ok {
			dedupKeys[item] = true
			result = append(result, item)
		}
	}
	return result
}
