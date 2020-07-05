package runner

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type TestSuites map[string]TestSuite
type TestSuite map[string]TestCase
type TestCase func(ctx context.Context, url string) error

func Suites(url string, suites TestSuites) {
	for suiteName, suite := range suites {
		fmt.Println()
		fmt.Println()
		fmt.Println("----------------------------------------------------")
		fmt.Printf("Test Suite: %s\n", suiteName)
		fmt.Println("----------------------------------------------------")
		Suite(url, suite)
	}
}

func Suite(url string, suite TestSuite) {
	testStartTime := time.Now()
	s := 0
	f := 0

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	results := make(TestResults, len(suite))

	for testID, testCase := range suite {
		fmt.Println("----------------------------------------------------")
		fmt.Printf("Case: %s\n", testID)
		fmt.Println("----------------------------------------------------")

		results.Start(testID)

		err := testCase(ctx, url)
		if err != nil {
			f++

			results.End(testID, false, err)
			takeFailureScreenshot(ctx, testID, err)
		} else {
			s++
			results.End(testID, true, nil)
		}

		fmt.Println("----------------------------------------------------")

		if !results.GetSuccess(testID) {
			fmt.Printf("Error   : %v\n", err)
		}

		fmt.Printf("Success : %v\n", results.GetSuccess(testID))
		fmt.Printf("Duration: %v\n", results.GetDuration(testID))
		fmt.Println("----------------------------------------------------")
	}

	fmt.Println("----------------------------------------------------")
	fmt.Printf("Duration: %v\n", time.Since(testStartTime))
	fmt.Printf("SUCCESS : %v\n", s)
	fmt.Printf("FAIL    : %v\n", f)

	results.GetFailed().PrintErrors()

	if len(results.GetFailed()) > 0 {
		os.Exit(1)
	}
}
