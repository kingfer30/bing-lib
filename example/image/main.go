package main

import (
	"fmt"
	"os"

	binglib "github.com/kingfer30/bing-lib"
)

var cookie = os.Getenv("COOKIE")

/*
生成图像
*/
func main() {
	i := binglib.NewImage(cookie)
	imgs, id, err := i.Image("猫") // 生成 4 张图片
	if err != nil {
		panic(err)
	}

	fmt.Println("id: ", id)
	fmt.Println("imgs: ", imgs)
}
