package util

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

//fmt.Sprintf("%v", util.IfThenElse(dev, "dev", "    "))

// IfThenElse evaluates a condition, if true returns the first string parameter otherwise the second
func IfThenStr(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}
