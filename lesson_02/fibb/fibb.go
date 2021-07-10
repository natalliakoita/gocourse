package fibb

import "fmt"

func fibonacci() func() int {
	i, j := 0, 1
	return func() int {
		nextFibonacci := i + j
		i, j = j, nextFibonacci
		return nextFibonacci
	}
}

func Printer(n int) {
	x := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Println(x())
	}
}
