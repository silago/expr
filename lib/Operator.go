package lib

import (
	"fmt"
	"math"
	"strings"
)

type Operator int

const (
	OpMul Operator = iota
	OpDiv
	OpAdd
	OpSub
)

var operatorMap = map[string]Operator{
	"+": OpAdd,
	"-": OpSub,
	"*": OpMul,
	"/": OpDiv,
	"=": OpDiv,
}



type OpTerm struct {
	Operator Operator `@("+" | "-" | "$")`
	Term     *Term    `@@`
}

type Term struct {
	Left  *Factor     `@@`
	Right []*OpFactor `{ @@ }`
}

type Factor struct {
	Base     *Value `@@`
	Exponent *Value `[ "^" @@ ]`
}

type OpFactor struct {
	Operator Operator `@("*" | "/")`
	Factor   *Factor  `@@`
}

func (o *Operator) Capture(s []string) error {
	*o = operatorMap[s[0]]
	return nil
}


func (o Operator) String() string {
	switch o {
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	case OpSub:
		return "-"
	case OpAdd:
		return "+"
	}
	panic("unsupported operator")
}


func (f *Factor) String() string {
	out := f.Base.String()
	if f.Exponent != nil {
		out += " ^ " + f.Exponent.String()
	}
	return out
}

func (f *Factor) Eval(ctx Context) float64 {
	b := f.Base.Eval(ctx)
	if f.Exponent != nil {
		return math.Pow(b, f.Exponent.Eval(ctx))
	}
	return b
}

// Display

func (o *OpFactor) String() string {
	return fmt.Sprintf("%s %s", o.Operator, o.Factor)
}

func (t *Term) String() string {
	out := []string{t.Left.String()}
	for _, r := range t.Right {
		out = append(out, r.String())
	}
	return strings.Join(out, " ")
}

func (o *OpTerm) String() string {
	return fmt.Sprintf("%s %s", o.Operator, o.Term)
}

func (o Operator) Eval(l, r float64) float64 {
	switch o {
	case OpMul:
		return l * r
	case OpDiv:
		return l / r
	case OpAdd:
		return l + r
	case OpSub:
		return l - r
	}
	panic("unsupported operator")
}





func (t *Term) Eval(ctx Context) float64 {
	n := t.Left.Eval(ctx)
	for _, r := range t.Right {
		n = r.Operator.Eval(n, r.Factor.Eval(ctx))
	}
	return n
}

