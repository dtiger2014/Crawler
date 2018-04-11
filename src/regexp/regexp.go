package main

import (
	"fmt"
	"regexp"
)

const (
	text = `
	My email is fanfan@qq.com
	email is abc@def.org
	email2 is      kkk@qq.com
	email3 is   ads@asb.com.cn
	`
)

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// match := re.FindAllString(text, -1)
	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}
}
