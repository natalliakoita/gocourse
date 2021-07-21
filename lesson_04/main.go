package main

import "github.com/natalliakoita/gocourse/lesson_04/arg"

func main() {
	c := arg.Circle{Radius: 6}
	r := arg.Rectangle{
		Height: 8,
		Width:  2,
	}
	err := arg.DescribeShape(c)
	if err != nil{
		panic(err)
	}
	err = arg.DescribeShape(r)
	if err != nil{
		panic(err)
	}
}
