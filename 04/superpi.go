package main

import (
	"fmt"
	"os"
	"strconv"
)

func main () {
    var input string
    fmt.Printf("请输入计算强度(1~100000000...):")
    _ ,err := fmt.Scanln(&input)
    if err != nil {
        fmt.Println("错误,请输入有效的计算强度")
        os.Exit(0)
    }
	num := 0.0
	ber := 1.0
	act := "add"
	in,err := strconv.ParseFloat(input,64)
	for ber <= float64(in) {
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
	fmt.Printf("圆周率的近似值计算得 %v\n",num*4)
}
