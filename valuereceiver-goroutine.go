package main

import (
	"fmt"
	"time"
)

func valueReceiverGoroutine() {
	// I used a value receiver to make sure the function (and communicate) that wouldn't mutate the struct, but it goes
	// both ways; mutations also aren't communicated to the function. This can lead to interesting bugs in concurrent
	// programming:
	subject := someStruct{value: "initial"}
	go subject.doWork()
	subject.value = "updated"
	time.Sleep(2 * time.Second)
}

type someStruct struct {
	value string
}

func (s someStruct) doWork() {
	fmt.Printf("doWork before sleep value=%s\n", s.value)
	time.Sleep(time.Second)
	fmt.Printf("doWork after sleep value=%s (expected: updated)\n", s.value)
}

