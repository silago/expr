package lib

import (
	"math"
	"strings"
)

type Func struct {
	Name  string     `@Ident`
	Right Expression `"(" { @@ } ")"`
}

func (f *Func) Eval(contexts Context) float64 {
	var result = f.Right.Eval(contexts)
	switch f.Name {
	case "floor":
		result = math.Floor(result)
	case "ceil":
		result = math.Ceil(result)
	}

	return result
}


func (f *Func) String() string {
	out := []string{f.Name}
	out = append(out, "(")
	out = append(out, f.Right.String())
	out = append(out, ")" )
	return strings.Join(out, " ")
}
