package main

import (
	"fmt"
	"os"

	binglib "github.com/kingfer30/bing-lib"
)

var cookie = os.Getenv("COOKIE")

/*
直接输出
*/
func main() {
	c := binglib.NewChat(cookie)
	c.NewConversation()

	r, err := c.Chat("", "你好")
	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}
