package interpreter

func is_int(i interface{}) bool {
	switch i.(type) {
	case int, uint:
		return true
	default:
		return false
	}
}
func is_float(i interface{}) bool {
	switch i.(type) {
	case float32, float64:
		return true
	default:
		return false
	}
}
func is_number(i interface{}) bool {
	switch i.(type) {
	case float32, float64, int, uint:
		return true
	default:
		return false
	}
}
