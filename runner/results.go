package runner

import (
	"fmt"
	"strconv"
	"time"
)

type TestResult struct {
	Success   bool
	StartTime time.Time
	EndTime   time.Time
	err       error
}

type TestResults map[string]TestResult

func (r TestResults) GetDuration(testID string) time.Duration {
	return r[testID].EndTime.Sub(r[testID].StartTime)
}

func (r TestResults) GetSuccess(testID string) bool {
	return r[testID].Success
}

func (r TestResults) Start(testID string) {
	r[testID] = TestResult{
		Success:   false,
		StartTime: time.Now(),
	}
}

func (r TestResults) End(testID string, v bool, err error) {
	result := r[testID]
	result.Success = v
	result.EndTime = time.Now()
	result.err = err
	r[testID] = result
}

func (r TestResults) PrintErrors() {
	if len(r) == 0 {
		return
	}

	fmt.Println("ERRORS:")

	maxLength := r.getMaxNameLength()

	for testName, result := range r {
		if result.err == nil {
			continue
		}

		fmt.Printf("%"+strconv.Itoa(maxLength)+"s: %v\n", testName, result.err)
	}
}

func (r TestResults) GetFailed() TestResults {
	failed := TestResults{}

	for testName, result := range r {
		if result.err == nil {
			continue
		}

		failed[testName] = result
	}

	return failed
}

func (r TestResults) getMaxNameLength() int {
	var max int

	for testName := range r {
		l := len(testName)
		if l > max {
			max = l
		}
	}

	return max
}
