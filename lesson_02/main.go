package main

import (
	"fmt"

	"github.com/natalliakoita/gocourse/lesson_02/fibb"
)

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Mission completed")
	fibb.Printer(5)
}
