package main

import (
	"fmt"
)

func countCollatzConjecTure(number int) {
	count := 0
	tempNumber := number
	for number != 1 {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = 3*number + 1
		}
		count = count + 1

	}
	fmt.Println("count CollatzConjec number ", tempNumber, "is: ", count)
}

func main() {
	var startNumber, endNumber int
	n, err := fmt.Scanln(&startNumber, &endNumber)
	_ = n
	if err != nil {
		fmt.Println("error: ", err)
	}
	for i := startNumber; i < endNumber; i++ {
		countCollatzConjecTure(i)
	}

}
