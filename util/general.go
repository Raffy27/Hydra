package util

//RemoveDuplicates returns a new slice with duplicate elements removed.
func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, i := range slice {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			list = append(list, i)
		}
	}

	return list
}
