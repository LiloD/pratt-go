package precedence

// define the precedence
const (
	Lowest = iota
	Conditional
	Or
	And
	Sum
	Product
	Exponent
	Prefix
	Call
)
