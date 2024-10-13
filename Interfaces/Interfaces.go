package main

import (
	"fmt"
	//"math"
)

type I interface{
	M()
}

type T struct {
	S string
}

func (t *T) M(){
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M(){
	fmt.Println(f)
}

func main(){
	// interface nil
	// var i I 

	// var t *T
	// i = t
	// describe(i)
	// i.M()

	// i = &T{"hello"}
	// describe(i)
	// i.M()

	//empty interface
	var j interface{}
	describe(j)

	j = 42
	describe(j)

	j = "hello"
	describe(j)
}

// func describe(i I){
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

func describe(j interface{}){
	fmt.Printf("(%v, %T)\n", j, j)
}

