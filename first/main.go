package main //without this go will not work

import (
	"fmt"
	"runtime"
)

func main() { //a special function that tells go where to start executing
	fmt.Println(runtime.NumCPU())
}
