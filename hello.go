package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i > 0; i++ {
		time.Sleep(5 * time.Second)
		fmt.Printf("Hello, world\n",i)
	}

}
