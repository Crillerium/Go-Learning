package main

import (
	"fmt"
)

func main () {
	num := 0.0
	ber := 1.0
	act := "add"
	for ber <= 1000000000 {
		if act == "add" {
			num += (1/ber)
			ber += 2
			act = "min"
		}else if act == "min" {
			num -= (1/ber)
			ber += 2
			act = "add"
		}
	}
	fmt.Printf("Pi may be %v\n",num*4)
}
