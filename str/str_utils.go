package str_utils

func IsEmpty(str string) bool {
	return !IsNotEmpty(str)
}

func IsNotEmpty(str string) bool {
	return &str != nil && len(str) > 0
}
