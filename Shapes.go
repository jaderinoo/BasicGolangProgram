package main

import "fmt"
import "encoding/json"

type Shape interface {
	Do()
}

type circ struct 
{
	Type int
	X, Y, R int
}

func (c *circ)Do(){
}

type rect struct {
	Type int
	X, Y, L, W int
}

func (c *rect)Do(){

}

type Drawing struct {
	Name   string
	Shapes []Shape
}

func main() {
	dr := Drawing{Name: "drawing a Circle with coordinates"}
	dr.Shapes = append(dr.Shapes, &circ{Type: 1, X: 2,Y: 3, R: 10})
	br := Drawing{Name: "drawing a Rectangle with coordinates"}
	br.Shapes = append(br.Shapes, &rect{Type: 1, X: 12,Y: 5, L: 5, W: 5,})
	
	jstr, err := json.Marshal(dr)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println(string(jstr))
	
	bstr, err := json.Marshal(br)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println(string(bstr))

}