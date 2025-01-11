package precedence

// define the precedence
const (
	Lowest = iota
	Sum
	Product
	Exponent
	Prefix
	Call
)
