package main

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestGetNextFriday(t *testing.T) {
	currentDate := time.Date(2023, 3, 18, 0, 0, 0, 0, time.UTC)

	expectedFriday := time.Date(2023, 3, 24, 0, 0, 0, 0, time.UTC)

	actualFriday := getNextFriday(currentDate)

	if actualFriday != expectedFriday {
		t.Errorf("Expected %v, got %v\n", expectedFriday, actualFriday)
	}
}

func TestPrintFridays(t *testing.T) {
	originalStdout := os.Stdout

	defer func() {
		os.Stdout = originalStdout
	}()

	testCases := []struct {
		startingFriday time.Time
		count          int
		expectedOutput string
	} {
		{
            startingFriday: time.Date(2022, time.January, 7, 0, 0, 0, 0, time.UTC),
            count:          3,
            expectedOutput:       "May 13, 2022\nJan 13, 2023\nOct 13, 2023\n",
        },
        {
            startingFriday: time.Date(2020, time.March, 6, 0, 0, 0, 0, time.UTC),
            count:          2,
            expectedOutput:       "Mar 13, 2020\nNov 13, 2020\n",
        },
	}

	for idx, cases := range testCases {
		r, w, _ := os.Pipe()
		os.Stdout = w

		printFridays(cases.startingFriday, cases.count)

		w.Close()

		actualOutput, _ := io.ReadAll(r)
		result := string(actualOutput)

		if result != cases.expectedOutput {
			t.Errorf("Test case %d: Expected %s, got %s\n", idx, cases.expectedOutput, result)
		}
	}
}