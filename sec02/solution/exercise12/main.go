// Section 02 : Lab 01 - Greeter
package main

import (
	"fmt"

	"github.com/egbordzor/learn-glft/shared/cal"
	"github.com/egbordzor/learn-glft/shared/input"
)

func main() {
	fmt.Print("What is your name: ")
	username := input.ReadString()
	printGreeting(username)
}

func printGreeting(username string) {
	fmt.Printf("\nHi, %s\n", username)
	fmt.Printf("Today is %s. The current time is %s.\n", cal.DateNow(), cal.TimeNow())
}
