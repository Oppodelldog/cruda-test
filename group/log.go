package group

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func Text(text string) chromedp.Action {
	return log{
		text: text,
	}
}

type log struct {
	text string
}

func (e log) Do(_ context.Context) error {
	fmt.Println(e.text)

	return nil
}
