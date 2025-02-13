package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is weekday")

	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It is boolean")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Printf("I don't know that type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
