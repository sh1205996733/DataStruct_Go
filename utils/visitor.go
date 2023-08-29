package utils

import "fmt"

type Visitor struct {
	Stop  bool
	Visit func(any any) bool
}

func NewVisitor() Visitor {
	return Visitor{
		Visit: func(val any) bool {
			fmt.Println(val)
			//fmt.Printf("%v\n", val)
			return false
		},
	}
}
