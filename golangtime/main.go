package main

import (
	"fmt"

	"time"
)

func main() {

	fmt.Println("Time study golang")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) //default formatting values.

}
