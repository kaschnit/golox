package conversion

// Return true if val is "truthy", otherwise return false.
func IsTruthy(val interface{}) bool {
	if val == nil {
		return false
	}

	switch t := val.(type) {
	case int:
		return t != 0
	case float32:
		return t != 0
	case float64:
		return t != 0
	case bool:
		return t
	case string:
		return t != ""
	default:
		// True for all non-null pointer types.
		return true
	}
}

func ToFloat(val interface{}) (float64, bool) {
	if val == nil {
		return 0, false
	}

	switch t := val.(type) {
	case bool:
		if t {
			return 1, true
		} else {
			return 0, true
		}
	case int:
		return float64(t), true
	case float32:
		return float64(t), true
	case float64:
		return t, true
	default:
		return 0, false
	}
}
