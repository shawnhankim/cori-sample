package main

import (
	"fmt"
)

type SampleState int

const (
	State00 SampleState = iota
	State01
	State02
)

// SampleStateName is the list of sample state's name
var SampleStateName = map[int]string{
	0: "State00",
	1: "State01",
	2: "State02",
}

// Rerun the string of sample state
func (st SampleState) String() string {
	names := [...]string{
		"State00",
		"State01",
		"State02"}
	if st < 0 || st > 2 {
		return "UNKNOWN"
	}
	return names[st]
}

func main() {
	fmt.Printf("Enum Value  : %d\n", State01)
	fmt.Printf("Enum String : %s\n", State01)
}
