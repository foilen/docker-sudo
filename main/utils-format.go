package main

func validateImageName(imageName string) bool {
	for _, r := range imageName {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == ':' || r == '.' || r == '-' || r == '_' || r == '/' {
			continue
		}
		return false
	}
	return true
}
