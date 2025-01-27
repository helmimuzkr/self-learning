package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for {
		fmt.Println(count)

		if count == 3 {
			var input string
			fmt.Scanf("%s", &input)
			fmt.Println(input)
		}

		count++

		time.Sleep(2 * time.Second)

	}
}
