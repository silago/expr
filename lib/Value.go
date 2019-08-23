package lib

import "fmt"

type Value struct {
	Number        *float64    `  @(Float|Int)`
	Func          *Func       `| @@ `
	Variable      *string     `| @Ident`
	Subexpression *Expression `| "(" { @@ } ")"`
}

func (v *Value) String() string {
	if v.Number != nil {
		return fmt.Sprintf("%g", *v.Number)
	}
	if v.Variable != nil {
		return *v.Variable
	}

	if v.Func != nil {
		return v.Func.String()
	}
	return "(" + v.Subexpression.String() + ")"
}

func (v *Value) GetVariableNames() (result []string) {
	if v.Number != nil {
	} else if v.Func != nil {
		result = v.Func.GetVariableNames()
	} else if v.Variable != nil {
		result = []string{*v.Variable}
	} else {
		result = v.Subexpression.GetVariableNames()
	}
	return
}

func (v *Value) Eval(ctx Context) float64 {
	switch {
	case v.Number != nil:
		return *v.Number
	case v.Variable != nil:
		value, ok := ctx[*v.Variable]
		if !ok {
			panic("no such variable " + *v.Variable)
		}
		return value
	case v.Func != nil:
		return v.Func.Eval(ctx)
		//panic("func")
	default:
		return v.Subexpression.Eval(ctx)
	}
}
