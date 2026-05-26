package kubeval

func getObject(body map[string]interface{}, key string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getStringAt(body map[string]interface{}, path []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getString(body map[string]interface{}, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// detectLineBreak returns the relevant platform specific line ending
func detectLineBreak(haystack []byte) string { _ = "STUB: not implemented"; return "" }

// in is a method which tests whether the `key` is in the set
func in(set []string, key string) bool { _ = "STUB: not implemented"; return false }
