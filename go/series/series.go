package series

const testVersion = 2

func All(n int, s string) []string {

	if n > len(s) {
		return nil
	}

	stringSeries := make([]string, len(s)-n+1)
	for start := 0; start+n <= len(s); start++ {
		stringSeries[start] = s[start : start+n]
	}
	return stringSeries
}

func First(n int, s string) string {
	if r := All(n, s); r != nil {
		return r[0]
	}
	return ""
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
