package utils

type Target string

func (t Target) In(x []string) bool {
	for _, val := range x {
		if t == Target(val) {
			return true
		}
	}
	return false
}
