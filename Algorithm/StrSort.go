package main

import (
	"fmt"
	"strings"
)

func StrSort(arr []string)[]string  {
	length:=len(arr)//取出数组长度
	if length<=1{
		return arr//单个元素直接返回
	} else{
		for i:=0;i<length-1;i++ { //只剩一个元素不需要挑选
			min:=i //标记索引
			for j:=i+1;j<length;j++{//每次选出一个极小值
				if strings.Compare(arr[min],arr[j])<0{
					min=j //保存极小值的索引
				}
			}
			if i!=min{
				arr[i],arr[min]=arr[min],arr[i]
			}
		}
		return arr
	}
}
func main()  {
	arr:=[]string{"c","a","s","d","f","e","z"}
	fmt.Println(StrSort(arr))
}
