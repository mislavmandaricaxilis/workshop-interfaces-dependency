package workshop_interfaces_dependency

import "fmt"

func NewLoggerOneWay() LoggerInterface {
	return loggerImplementation{}
}

type LoggerInterface interface {
	Log(s string)
}

type loggerImplementation struct {}

func (l loggerImplementation) Log(s string) {
	fmt.Println(s)
}