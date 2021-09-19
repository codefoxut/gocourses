package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(e *DoubleExpression3)
	VisitAdditionExpression(a *AdditionalExpression3)
}

type Expression3 interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression3 struct {
	value float64
}

func (d *DoubleExpression3) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(d)
}

type AdditionalExpression3 struct {
	left, right Expression3
}

func (a *AdditionalExpression3) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(a)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{strings.Builder{}}
}

func (ep *ExpressionPrinter) VisitDoubleExpression(d *DoubleExpression3) {
	ep.sb.WriteString(fmt.Sprintf("%g", d.value))
}

func (ep *ExpressionPrinter) VisitAdditionExpression(a *AdditionalExpression3) {
	ep.sb.WriteRune('(')
	a.left.Accept(ep)
	ep.sb.WriteRune('+')
	a.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func NewExpressionEvaluator() *ExpressionEvaluator {
	return &ExpressionEvaluator{}
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(d *DoubleExpression3) {
	ee.result = d.value
}

func (ee *ExpressionEvaluator) VisitAdditionExpression(a *AdditionalExpression3) {
	a.left.Accept(ee)
	x := ee.result
	a.right.Accept(ee)
	x += ee.result
	ee.result = x
}

func ExecuteVisitorClassic() {
	fmt.Println("classic")
	e := &AdditionalExpression3{
		left: &DoubleExpression3{1},
		right: &AdditionalExpression3{
			left:  &DoubleExpression3{2},
			right: &DoubleExpression3{3},
		},
	}
	ep := NewExpressionPrinter()
	e.Accept(ep)

	ee := NewExpressionEvaluator()
	e.Accept(ee)
	fmt.Println(ep.String(), "=", ee.result)

}
