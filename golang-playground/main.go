package main

import (
	"fmt"
	"time"
)

func main() {
	// Format uses a layout defined as a default date 2 Jan 2006 which will be used to format the time,
	// let’s see a couple of examples to better understand how this works.
	fmt.Println(time.Now().Format("2006-01-02;15:04:05"))
}
