package main

import (
	"DataStruct_Go/print/printer"
)

func main() {
	root := &Person{"B",
		&Person{"A",
			&Person{"A1", nil, nil, false},
			&Person{"A2", nil, nil, true}, false},
		&Person{"C",
			&Person{"C1", nil, nil, false},
			&Person{"C2", nil, nil, false}, false}, false}

	printer.Println(root)
}

type Person struct {
	Value string
	Left  *Person
	Right *Person
	Color bool
}

func (p *Person) RootNode() interface{} {
	return p
}
func (p *Person) LNode() interface{} {
	return p.Left
}
func (p *Person) RNode() interface{} {
	return p.Right
}
func (p *Person) ToString() interface{} {
	return p.Value
}
func (p *Person) ColorOf() bool {
	return p.Color
}
