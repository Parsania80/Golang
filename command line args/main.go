package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path: ", os.Args[0])
	// fmt.Println("first  arg: ", os.Args[1])
	// fmt.Println("second arg: ", os.Args[2])

	result, err := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	_ = err

}
