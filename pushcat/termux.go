package main

import (
    "fmt"
	"flag"
	"os"
	"os/exec"
	"log"
)

func main () {
	var title string
	var content string
	flag.StringVar(&title,"t","默认の标题","标题")
	flag.StringVar(&content,"c","未定义的内容","内容")
	flag.Parse()
	cmd := exec.Command("termux-notification", "-t", title, "-c", content)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(0)
	}
	fmt.Println("推送完成")
}
