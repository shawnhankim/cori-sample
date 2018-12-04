// testTimer.go
package main

import (
    "runtime"
    "time"
)

func retimer(t *time.Timer, d time.Duration) {
    if !t.Stop() {
        select {
        case <-t.C:
        default:
        }
    }
    t.Reset(d)
}

func runTimer() {
    tmr := time.NewTimer(0)
    retimer(tmr, time.Minute)
    select {
    case <-tmr.C:
        panic("unexpected firing of Timer")
    default:
    }
}

func main() {
    runtime.GOMAXPROCS(2)
    for {
        runTimer()
    }
}
