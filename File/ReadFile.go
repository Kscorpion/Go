package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("嗨客网(www.haicoder.net)")
	fileName := "1.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Read File err, err =", err)
		return
	}
	defer file.Close()
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		//说明读取结束
		if n == 0 {
			break
		}
		fmt.Println(string(buf))
		//读取到最终的缓冲区中
		//chunk = append(chunk, buf[:n]...)
	}
	fmt.Println("File Content =", string(chunk))
}
