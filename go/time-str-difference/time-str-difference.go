package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time parsing")
	dateStr1 := "2018-11-28T17:38:45Z"
	dateStr2 := "2018-11-28T16:38:45Z"
	tCur := time.Now()

	t1, err1 := time.Parse(time.RFC3339, dateStr1)
	if err1 != nil {
		fmt.Println("Error while parsing date :", err1)
	}
	t2, err2 := time.Parse(time.RFC3339, dateStr2)
	if err2 != nil {
		fmt.Println("Error while parsing date :", err2)
	}
	fmt.Printf("The hours   diffrence is: %f \n", t1.Sub(t2).Hours())
	fmt.Printf("The seconds diffrence is: %f \n", t1.Sub(t2).Seconds())
	fmt.Printf("The time    diffrence is: %f \n", tCur.Sub(t2).Seconds())

}
