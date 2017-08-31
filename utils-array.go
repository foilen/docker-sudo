package main

func arrayContains(array []string, val string) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}

	return false
}
