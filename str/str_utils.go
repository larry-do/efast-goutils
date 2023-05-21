package str_utils

func IsEmpty(str string) bool {
	if &str == nil || len(str) > 0 {
		return false
	}
	return true
}
