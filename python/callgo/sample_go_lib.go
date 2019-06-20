package main

import "C"

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

var count int

var mtx sync.Mutex

//export Add
func Add(a, b int) int { return a + b }

//export Cosine
func Cosine(x float64) float64 { return math.Cos(x) }

//export Sort
func Sort(vals []int) { sort.Ints(vals) }

//export Log
func Log(msg string) int {
	mtx.Lock()
	defer mtx.Unlock()
	strMsg := fmt.Sprintf("Golang logging message : %s", msg)
	fmt.Println(strMsg)
	count++
	return count
}

func main() {}
