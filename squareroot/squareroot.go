package main

import (
	"fmt"
	"math"
	"strconv"
)

func squareroot(num float64) float64 {
	return math.Sqrt(num)
}

func main () {
	var input string
	for true {
		fmt.Printf("请输入要开平方的数字:")
		_ ,err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("错误,请输入一个有效的数字!")
			continue
		}
		in,err := strconv.ParseFloat(input,64)
		fmt.Printf("%v 的平方根计算得 %v\n",input,squareroot(in))
	}
}
