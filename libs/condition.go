package libs

import (
    "fmt"
    "time"
)

// https://www.w3schools.com/go/go_conditions.php

func IsBiggerThan10(num int) bool {
	var rtn bool
	if num > 10 {
		rtn = true
	} else {
		rtn = false
	}
	return rtn
}

func SayGreetingForCurrectTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
