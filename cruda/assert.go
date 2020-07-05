package cruda

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func Equals(expression string, expected interface{}) chromedp.Action {
	return prefixError("Equals", EqualsAction{
		Expression: expression,
		Expected:   expected,
	})
}

type EqualsAction struct {
	Expression string
	Expected   interface{}
}

func (e EqualsAction) Do(ctx context.Context) error {
	var res interface{}

	err := chromedp.Evaluate(e.Expression, &res, chromedp.EvalAsValue).Do(ctx)
	if err != nil {
		fmt.Printf("evaluation error: %v", err)
		return err
	}

	isEqual := false

	switch v1 := e.Expected.(type) {
	case int:
		v2, ok := res.(int)
		if !ok {
			v2f, okf := res.(float64)
			if okf {
				v2 = int(v2f)
				ok = true
			}
		}

		isEqual = ok && v1 == v2
	case string:
		v2, ok := res.(string)
		isEqual = ok && v1 == v2
	case bool:
		v2, ok := res.(bool)
		isEqual = ok && v1 == v2
	}

	if !isEqual {
		return fmt.Errorf("expected type was >>%v<< (%T), but chrome returned >>%v<< (%T)", e.Expected, e.Expected, res, res)
	}

	return nil
}
