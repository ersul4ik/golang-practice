package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func main() {
	var fNum, sNum int

	_, err := fmt.Scan(&fNum, &sNum)
	if err != nil {
		fmt.Print(err)
		return
	}

	multiply := multiply(fNum, sNum)

	fmt.Printf("Input Numbers Multiplication is: %d", multiply)
}
