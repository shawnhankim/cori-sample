package main

import (
    "fmt"
    "time"
)

func GetCurTimeStr() string {
    t := time.Now();
    return t.Format("2006-01-02T15:04:05");
}


func main() {
    fmt.Printf("%s\n", GetCurTimeStr());
}

