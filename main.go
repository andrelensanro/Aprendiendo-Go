package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, -1}
	p := &v
	p.X = -2
	fmt.Println(v)

	v1 := Vertex{X:-4}
	v3 := Vertex{Y:7}

	fmt.Println(v1, v3)

	s := []struct{
		i int 
		b bool
		s string
	}{
		{1, true, "Hello"},
		{2, false, "World"},
		{3, true, "Holaaaa"}}
	
	fmt.Println(s)

	var arr []int
	if arr == nil {
		fmt.Println("nil")
	}
}