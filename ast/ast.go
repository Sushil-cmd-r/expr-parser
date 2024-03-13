package ast

import (
	"strconv"

	"github.com/sushil-cmd-r/expr-parser/token"
)

type Expr interface {
	exprNode()
}

type BinaryExpr struct {
	Operator token.Token
	Left     Expr
	Right    Expr
}

func NewBinary(Operator token.Token, Left, Right Expr) BinaryExpr {
	return BinaryExpr{
		Operator: Operator,
		Left:     Left,
		Right:    Right,
	}
}

type UnaryExpr struct {
	Operator token.Token
	Right    Expr
}

func NewUnary(Operator token.Token, Right Expr) UnaryExpr {
	return UnaryExpr{
		Operator: Operator,
		Right:    Right,
	}
}

type NumberExpr struct {
	Value float64
}

func NewNumber(Value string) NumberExpr {
	num, err := strconv.ParseFloat(Value, 64)
	if err != nil {
		return NumberExpr{}
	}

	return NumberExpr{
		Value: num,
	}
}

func (b BinaryExpr) exprNode() {}
func (u UnaryExpr) exprNode()  {}
func (n NumberExpr) exprNode() {}
