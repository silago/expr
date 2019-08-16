package expr
/*
	based on
	https://github.com/alecthomas/participle/blob/master/_examples/expr/main.go
 */
import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)



type Expression struct {
	Left  *Term     `@@`
	Right []*OpTerm `{ @@ }`
}

func (e *Expression) String() string {
	out := []string{e.Left.String()}
	for _, r := range e.Right {
		out = append(out, r.String())
	}
	return strings.Join(out, " ")
}

func (e *Expression) Eval(ctx Context) float64 {
	l := e.Left.Eval(ctx)
	for _, r := range e.Right {
		l = r.Operator.Eval(l, r.Term.Eval(ctx))
	}
	return l
}

type Context map[string]float64

var convertonError = errors.New("invalid context value")
func any2float64 (v interface{}) ( float64, error ) {
	switch v.(type) {
		case int:
			return float64(v.(int)), nil
		case int64:
			return float64(v.(int64)), nil
		case float64:
			return float64(v.(float64)), nil
		case float32:
			return float64(v.(float32)), nil
		case string:
			if v, e:=strconv.ParseFloat(v.(string),64);e==nil {
				return v, nil
			} else {
				return 0, convertonError
			}
		default:
			var str = fmt.Sprintf("%v", v)
			if v, e:=strconv.ParseFloat(str,64);e==nil {
				return v, nil
			} else {
				return 0, convertonError
			}
	}
}


func prepareCtx(ctx map[string]interface{}) (Context, error) {
	var resultCtx = make(Context)
	var err error
	for k, v:= range ctx {
		if resultCtx[k], err = any2float64(v); err!=nil {
			return resultCtx, err
		}
	}
	return resultCtx, err
}

func ParseExpresstion(str string, inputContext map[string]interface{}) ( result float64, err error ) {
	var ctx Context
	if inputContext == nil {
		inputContext = map[string]interface{}{}
	} else if ctx, err = prepareCtx(inputContext); err !=nil {
		return 0, err
	}

	expr := &Expression{}

	if parser, err := participle.Build(&Expression{}); err!=nil {
		return 0, err
	} else if err:= parser.ParseString(str, expr); err != nil {
		return 0, err
	} else {
		var result = expr.Eval(ctx)
		return result, nil
	}
}
