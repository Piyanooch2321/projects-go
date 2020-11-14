package main

import (
	"fmt"
)

func main() {
	var array [5]int
	for i := range array {
		fmt.Println(array[i])
	}
}
