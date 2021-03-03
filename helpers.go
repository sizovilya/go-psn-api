package psn

func isContain(arr []string, value string) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}
	return false
}
