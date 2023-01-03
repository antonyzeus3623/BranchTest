package main

import "fmt"

type People struct {
	Name string
}

func (p *People) _String() string {
	return fmt.Sprintf("print: %v", p)
}
func main() {
	p := &People{}
	p._String()
}
