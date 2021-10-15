package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Now \t:%v\n", time.Now())
	fmt.Printf("Weekday \t:%v\n", time.Now().Weekday())
	fmt.Println(time.Now().Clock())
	fmt.Printf("Hour \t:%v\n", time.Now().Hour())
	fmt.Printf("Minute \t:%v\n", time.Now().Minute())
	fmt.Println(time.Now().Date())
	fmt.Printf("Month \t:%v\n", time.Now().Month())
	fmt.Printf("Year \t:%v\n", time.Now().Year())
}
