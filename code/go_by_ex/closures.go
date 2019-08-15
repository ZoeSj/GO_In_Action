package main

import "fmt"

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
