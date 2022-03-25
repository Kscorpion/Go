package main

type MessageV1 interface {
	Reset()
	String() string
	ProtoMessage() int
}

type XXX struct{}
type YYY struct{}

func main() {
	x := &XXX{}
	y := &YYY{}
	abcTest(x)
	abcTest(y)
}

//x 和 y有共同实现的接口
//都可以传入 abcTest()
func abcTest(data MessageV1) {

}

func (x *XXX) Reset() {

}

func (x *XXX) String() string {
	return ""
}

func (x *XXX) ProtoMessage() int {
	return 0
}

func (y *YYY) Reset() {

}

func (y *YYY) String() string {
	return ""
}

func (y *YYY) ProtoMessage() int {
	return 0
}
