package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func main() {
	fmt.Println("main", GoID())
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, GoID())
		}()
	}
	fmt.Println("++++++++++++++")
	for i := 1 << 10; i > 0; i-- {
		i++
		fmt.Print("/")
		i--
	}
	fmt.Println("")
	fmt.Println("++++++++++++++")
	wg.Wait()
}
