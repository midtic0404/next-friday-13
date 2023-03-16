package main

import (
	"fmt"
	"time"
)

const numberOfFridays = 10

func main() {
	todayDate := time.Now()
	startingFriday := getNextFriday(todayDate)

	fmt.Println("Starting Friday: ", startingFriday.Format("2006-01-02"))

	printFridays(startingFriday, numberOfFridays)
}

func getNextFriday(date time.Time) time.Time {
	for {
		date = date.AddDate(0, 0, 1)
		if date.Weekday() == time.Friday {
			return date
		}
	}
}

func printFridays(startingFriday time.Time, count int) {
	counter := 0
	for counter < count {
		if startingFriday.Weekday() == time.Friday && startingFriday.Day() == 13 {
			fmt.Println(startingFriday.Format("Jan 2, 2006"))
			counter++
		}

		startingFriday = startingFriday.AddDate(0, 0, 7)
	}
}