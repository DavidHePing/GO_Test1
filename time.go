package main

import (
	"fmt"
	"time"
)

func time_test1() {
	var time1 time.Time

	fmt.Println("check if time1 is default:", time1.IsZero())

	time1 = time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)

	fmt.Println("check if time1 is default:", time1.IsZero())
	fmt.Println(time1)
}
