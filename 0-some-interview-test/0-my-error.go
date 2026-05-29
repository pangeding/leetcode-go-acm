package main

import (
	"fmt"
)

type MyError struct {
	msg string
}

/*
*

	func Error(e *MyError) string {
		return e.msg
	}
*/
func (e *MyError) Error() string {
	return e.msg
}

func getError() error {
	var err *MyError = nil
	return err
}

/**
类型：*main.MyError
值：<nil>
是否为nil：false
错误信息： 出错了
*/
func main() {
	err := getError()
	fmt.Printf("类型：%T\n", err)
	fmt.Printf("值：%v\n", err)
	fmt.Printf("是否为nil：%v\n", err == nil)

	err2 := &MyError{msg: "出错了"}
	fmt.Printf("错误信息： %s\n", err2)
}
