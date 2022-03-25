package main

import (
	"fmt"
	"time"
)

//获取距离每日24时秒数
func main() {
	now := time.Now()
	fmt.Println(time.Until(time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local)).Seconds())
}
