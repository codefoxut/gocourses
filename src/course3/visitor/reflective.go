package visitor

import (
	"fmt"
	"strings"
)

type Expression2 interface {
}

type DoubleExpression2 struct {
	value float64
}

type AdditionalExpression2 struct {
	left, right Expression2
}

func Print(e Expression2, sb *strings.Builder) {
	if de, ok := e.(*DoubleExpression2); ok {
		sb.WriteString(fmt.Sprintf("%g", de.value))
	} else if ae, ok := e.(*AdditionalExpression2); ok {
		sb.WriteRune('(')
		Print(ae.left, sb)
		sb.WriteRune('+')
		Print(ae.right, sb)
		sb.WriteRune(')')
	}
}

func ExecuteVisitorReflective() {
	fmt.Println("reflective")
	e := &AdditionalExpression2{
		left: &DoubleExpression2{1},
		right: &AdditionalExpression2{
			left:  &DoubleExpression2{2},
			right: &DoubleExpression2{3},
		},
	}

	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}
