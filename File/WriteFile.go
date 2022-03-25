package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("1.txt", os.O_WRONLY|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i := 100001; i < 10000000; i++ {
			f.Write([]byte(fmt.Sprintf("%v 要写入的文本内容,time:%v\n", i, time.Now())))
		}
	}
}
