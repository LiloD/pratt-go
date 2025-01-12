package precedence

// define the precedence
const (
	Lowest = iota
	Conditional
	Sum
	Product
	Exponent
	Prefix
	Call
)
