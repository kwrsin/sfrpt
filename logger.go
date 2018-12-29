// logger.go
package main

import (
    "log"
    "os"
)

var (
    l  = log.New(os.Stderr, "", log.LstdFlags)
    Print   = l.Print
    Printf  = l.Printf
    Println = l.Println
    Panic   = l.Panic
    Panicf  = l.Panicf
    Panicln = l.Panicln
    Fatal   = l.Fatal
    Fatalf  = l.Fatalf
    Fatalln = l.Fatalln
)
