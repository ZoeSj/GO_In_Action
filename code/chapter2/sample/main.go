package main

import (
	"log"
	"os"
	//这里的下划线是为了让GO语言对包做初始化操作，但是并不使用包里的标识符。这样做的目的是调用matchers包中的rss.go中init函数。
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
//会在main函数执行前调用
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
