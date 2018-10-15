package util

func Unset(ss interface{}, i int) interface{} {
	s := ss.([]interface{})
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
