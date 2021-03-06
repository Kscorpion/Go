package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type RedPackage struct {
	RemainSize  float64 //剩余红包数量
	RemainMoney float64 //剩余金额
}

// 抢红包
func (rp *RedPackage) GrabRedPackage() float64 {
	if rp.RemainSize == 1 {
		rp.RemainSize--
		return float64(math.Round(rp.RemainMoney*100) / 100)
	}
	r := rand.Float64()
	min := 0.01
	max := rp.RemainMoney / rp.RemainSize * 2
	money := r * max
	if money <= min {
		money = 0.01
	}
	money = math.Floor(money*100) / 100
	rp.RemainSize--
	rp.RemainMoney -= money
	return money
}

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			r := RedPackage{float64(10), float64(30)}
			fmt.Println(r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage(), r.GrabRedPackage())
		}()
	}
	time.Sleep(time.Second * 2)
}
