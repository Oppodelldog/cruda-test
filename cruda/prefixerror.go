package cruda

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func prefixError(prefix string, action chromedp.Action) chromedp.Action {
	return errPrefixAction{
		Prefix: prefix,
		Action: action,
	}
}

type errPrefixAction struct {
	Prefix string
	Action chromedp.Action
}

func (e errPrefixAction) Do(ctx context.Context) error {
	err := e.Action.Do(ctx)
	if err != nil {
		return fmt.Errorf("%s: %v", e.Prefix, err)
	}

	return nil
}
