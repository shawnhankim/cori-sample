package main

import (
	"fmt"
)

type SampleState int

const (
	State00 SampleState = iota
	State01
	State02
	State03
)

// SampleStateName is the list of sample state's name
var SampleStateName = map[int]string{
	0: "StateName00",
	1: "StateName01",
	2: "StateName02",
	3: "StateName03",
}

// Rerun the string of sample state
func (st SampleState) String() string {
	names := [...]string{
		"StateName00",
		"StateName01",
		"StateName02",
		"StateName03"}
	if st < State00 || st > State03 {
		return "UNKNOWN"
	}
	return names[st]
}

func main() {
	fmt.Printf("Enum Value  : %d\n", State01)
	fmt.Printf("Enum String : %s\n", State01)
}
