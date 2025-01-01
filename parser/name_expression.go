package parser

import (
	"fmt"
)

// name
type NameExpression struct {
	Name string
}

// name should implement Expression
func (n *NameExpression) String() string {
	return fmt.Sprintf("(NAME: %s)", n.Name)
}

func (n *NameExpression) expression() {}
