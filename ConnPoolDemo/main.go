package main

import "ConnPoolDemo/ConnPool"

func main() {
	pool := ConnPool.NewPool(1, 5)
	for i := 0; i < 10; i++ {
		//获取链接
		p := pool.Get()

		//将链接放回
		pool.Put(p)
	}

}
