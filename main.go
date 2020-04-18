package main

import "fmt"

func multiplication(a int, b int) int {
	return a * b
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	var fNum, sNum int

	_, err := fmt.Scan(&fNum, &sNum)
	if err != nil {
		fmt.Print(err)
		return
	}

	result := multiplication(fNum, sNum)
	fmt.Printf("Input Numbers Multiplication is: %d\n", result)

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
}
