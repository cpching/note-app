package main

import (
    "sync"
)

var (
    timeLayout = "2006-01-02 15:04:05" // time parse format
    mu sync.Mutex
)

