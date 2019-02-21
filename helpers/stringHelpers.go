package helper

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

//Comparable takes a generic type and compares a and b
func Comparable(a interface{}, b interface{}) bool {
	return a == b
}
