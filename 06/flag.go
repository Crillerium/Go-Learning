package main

import (
	"fmt"
	"flag"
)

func main () {
	var title string
	var content string
	flag.StringVar(&title,"t","默认の标题","标题")
	flag.StringVar(&content,"c","黑犬认~の内容","内容")
	flag.Parse()
	fmt.Printf("我们的标题是:%s!\n",title)
	fmt.Printf("我们的内容是:%s!\n",content)
}
