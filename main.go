package main

import (
	"fmt"
	"time"
)

func main() {
	todayDate := time.Now()
	dayOfWeekNumber := int(todayDate.Weekday())
	
	daysToFriday := 5 - dayOfWeekNumber

	if daysToFriday < 0 {
		daysToFriday += 7
	}

	startingFriday := todayDate.AddDate(0, 0, daysToFriday)

	fmt.Println("Starting Friday: ", startingFriday.Format("2006-01-02"))

	printFridays(startingFriday, 10)
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