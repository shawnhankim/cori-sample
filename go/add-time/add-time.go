package main

import (
    "time"
    "fmt"
)

func main() {
    curTimeStr  := time.Now().Format("2006-01-02T15:04:05Z") 
    curTime,err := time.Parse(time.RFC3339, curTimeStr)
    if err != nil {
        fmt.Printf("Error : %v", err)
    }
    expTime     := curTime.AddDate(0,0,1)

    fmt.Printf("Cur Time : %s\n", curTimeStr)
    fmt.Printf("Exp Time : %s\n", expTime.Format("2006-01-02T15:04:05Z"))
}

