package workshop_interfaces_dependency

import "fmt"

func NewLoggerOtherWay() LoggerStruct {
	return LoggerStruct{}
}

type LoggerStruct struct {}

func (l LoggerStruct) Log(s string, data string) {
	fmt.Println(s)
	fmt.Println(data)
}