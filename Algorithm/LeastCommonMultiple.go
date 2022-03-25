package main

import (
	"math"
	"fmt"
)

type LeastCommonMultiple struct{}

func (lgb *LeastCommonMultiple) ret(a int ,b int )(c int) {
	var(
		max int
		min int
		ret int
	)
	max = int(math.Max(float64(a),float64(b)))
	min = int(math.Min(float64(a),float64(b)))
	for{
		ret = max%min
		if(ret == 0){//此时最大公约数为min，最小公倍数为两数相乘除以最大公约数
			return a*b/min
		}
		max = min
		min = ret
	}
}

func main(){
	lcm:=&LeastCommonMultiple{}
	r:=lcm.ret(3,4)
	fmt.Println(r)
}
