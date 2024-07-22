package main

import (
	"fmt"
	"math"
)

func squareroot(num float64) float64 {
	return math.Sqrt(num)
}

func main () {
	var input float64
	for true {
		fmt.Printf("请输入要开平方的数字:")
		_ ,err := fmt.Scanf("%f",&input)
		if err != nil {
			fmt.Println("错误,请输入一个有效数字!")
			continue
		}
		fmt.Printf("%v 的平方根为 %v\n",input,squareroot(input))
	}
}
