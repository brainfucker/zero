package zero

// H is general type of any free structured data
// for example using H you can describe custom json easily
type H map[string]interface{}

// String return field as string
func (h H) String(field string) string {
	val, ok := h[field]
	if ok {
		return J(val)
	}
	return ""
}

// Int return field as int
func (h H) Int(field string) int {
	val, ok := h[field]
	if ok {
		res, ok := val.(int)
		if ok {
			return res
		}
	}
	return 0
}

// Int64 return field as int64
func (h H) Int64(field string) int64 {
	val, ok := h[field]
	if ok {
		res, ok := val.(int64)
		if ok {
			return res
		}
	}
	return int64(0)
}

// Float return field as float64
func (h H) Float(field string) float64 {
	val, ok := h[field]
	if ok {
		res, ok := val.(float64)
		if ok {
			return res
		}
	}
	return float64(0)
}

// H return field as inner object of H
func (h H) H(field string) H {
	val, ok := h[field]
	if ok {
		res, ok := val.(H)
		if ok {
			return res
		}
	}
	return H{}
}

// IsEmpty will check
func (h H) IsEmpty() bool {
	return len(h) == 0
}
