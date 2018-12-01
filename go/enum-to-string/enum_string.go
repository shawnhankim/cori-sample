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
	0: "StateName00",
	1: "StateName01",
	2: "StateName02",
}

// Rerun the string of sample state
func (st SampleState) String() string {
	names := [...]string{
		"State00",
		"State01",
		"State02"}
	if st < State00 || st > State02 {
		return "UNKNOWN"
	}
	return names[st]
}

func main() {
	fmt.Printf("Enum Value  : %d\n", State01)
	fmt.Printf("Enum String : %s\n", State01)
}
