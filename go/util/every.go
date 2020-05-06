package util

func Every(vs []bool) bool {
	for _, v := range vs {
		if !v {
			return false
		}
	}
	return true
}