# Friday the 13th Finder
This is a simple command-line program that uses the Go programming language to find and print the next 10 Friday the 13th dates, starting from the current date. The program calculates the next Friday and then iterates through the following Fridays, checking if each one is the 13th day of the month.


## How to use
- Run go build
``` 
go build 
```

- Run program 
```
./next-friday-13
```

## Example

Suppose today's date is March 18, 2023 (Saturday). Running the program will first find the next Friday and then print the next `numberOfFridays` Friday the 13ths. 

For example, if we set `numberOfFridays` to 5, the program will output:



Starting Friday: Mar 24, 2023
Oct 13, 2023
Sep 13, 2024
Dec 13, 2024
Jun 13, 2025
Feb 13, 2026


In this example, the program identifies the next Friday as March 24, 2023, and then proceeds to print the next 5 Friday the 13ths.

