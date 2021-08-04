package main

import "fmt"

func InsertSort(arr []int)[]int {
	length:=len(arr)//取出数组长度
	if length<=1{
		return arr//单个元素直接返回
	}else{
		for i:=1;i<length;i++{//跳过第一个
			backup := arr[i]//备份插入的数据
			j:=i-1//上一个位置循环找到位置插入
			for j>=0 && backup<arr[j] {
				arr[j+1]=arr[j]//从前往后移动
				j--
			}
			arr[j+1]=backup

		}
		return arr
	}
}
func main(){
	arr:=[]int{33,5,6,2,3,8,0,4}
	fmt.Println(InsertSort(arr))
}
